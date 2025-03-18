package backoff_chain

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"
)

type testTimer struct {
	timer *time.Timer
}

func (t *testTimer) Start(duration time.Duration) {
	t.timer = time.NewTimer(0)
}

func (t *testTimer) Stop() {
	if t.timer != nil {
		t.timer.Stop()
	}
}

func (t *testTimer) C() <-chan time.Time {
	return t.timer.C
}

func TestRetryWithData(t *testing.T) {
	const successOn = 3
	var i = 0

	// This function is successful on "successOn" calls.
	f := func() (any, error) {
		i++
		log.Printf("function is called %d. time\n", i)

		if i == successOn {
			log.Println("OK")
			return 42, nil
		}

		log.Println("error")
		return 1, errors.New("error")
	}

	receiver := 0

	err := NewExponentialBackOff().WithInitialInterval(1*time.Second).WithMaxTries(4).WithReceiver(&receiver).Retry(context.Background(), f)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	if i != successOn {
		t.Errorf("invalid number of retries: %d", i)
	}
	if receiver != 42 {
		t.Errorf("invalid data in response: %d, expected 42", receiver)
	}
	t.Logf("res: %d", receiver)

}
