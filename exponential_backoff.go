package backoff_chain

import (
	"time"

	"github.com/cenkalti/backoff/v5"
)

type exponentialBackOff struct {
	*backoffDoer
	*backoff.ExponentialBackOff
}

// NewExponentialBackOff creates an instance of exponentialBackOff using default values.
func NewExponentialBackOff() *exponentialBackOff {
	origin := backoff.NewExponentialBackOff()
	eb := &exponentialBackOff{
		ExponentialBackOff: origin,
	}
	eb.backoffDoer = CustomizeBackoff(eb)
	return eb
}

func (eb *exponentialBackOff) WithInitialInterval(d time.Duration) *exponentialBackOff {
	eb.InitialInterval = d
	return eb
}

func (eb *exponentialBackOff) WithRandomizationFactor(f float64) *exponentialBackOff {
	eb.RandomizationFactor = f
	return eb
}

func (eb *exponentialBackOff) WithMultiplier(f float64) *exponentialBackOff {
	eb.Multiplier = f
	return eb
}

func (eb *exponentialBackOff) WithMaxInterval(d time.Duration) *exponentialBackOff {
	eb.MaxInterval = d
	return eb
}
