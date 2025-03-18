package backoff_chain

import (
	"context"
	"fmt"
	"reflect"
	"time"

	backoff "github.com/cenkalti/backoff/v5"
)

func CustomizeBackoff(b backoff.BackOff) *backoffDoer {
	pl := &backoffDoer{
		backoffBase: &backoffBase{
			maxElapsedTime: backoff.DefaultMaxElapsedTime,
		},
		b: b,
	}
	pl.backoffBase.payload = pl
	return pl
}

func (pl *backoffDoer) Retry(ctx context.Context, operation Operation) error {
	opts := make([]backoff.RetryOption, 0)
	if pl.b == nil {
		return fmt.Errorf("missing backoff parameter")
	}
	if pl.notify != nil {
		opts = append(opts, backoff.WithNotify(func(err error, duration time.Duration) {
			pl.notify(err, duration)
		}))
	}
	if pl.maxTries > 0 {
		opts = append(opts, backoff.WithMaxTries(pl.maxTries))
	}
	if pl.maxElapsedTime > 0 {
		opts = append(opts, backoff.WithMaxElapsedTime(pl.maxElapsedTime))
	}

	anyResp, err := backoff.Retry[any](ctx, func() (any, error) {
		return operation()
	}, opts...)
	if err != nil {
		return err
	}
	//save anyResp to result

	if pl.result != nil {
		// 使用反射获取 Result 的实际值和类型
		receiverVal := reflect.ValueOf(pl.backoffBase.result)
		// 确保 Result 是一个指针
		if receiverVal.Kind() != reflect.Ptr {
			return fmt.Errorf("receiver must be a pointer")
		}
		// 获取指针指向的元素
		elemVal := receiverVal.Elem()
		// 如果元素可设置
		if elemVal.CanSet() {
			// 获取 anyResp 的反射值
			respVal := reflect.ValueOf(anyResp)
			// 检查类型兼容性
			if respVal.Type().AssignableTo(elemVal.Type()) {
				// 设置值
				elemVal.Set(respVal)
			} else {
				return fmt.Errorf("operation result type %v cannot be assigned to receiver type %v",
					respVal.Type(), elemVal.Type())
			}
		} else {
			return fmt.Errorf("receiver is not settable")
		}
	}
	return nil
}
