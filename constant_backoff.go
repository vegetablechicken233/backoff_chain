package backoff_chain

import (
	"time"

	backoff "github.com/cenkalti/backoff/v5"
)

type constantBackOff struct {
	*backoffDoer
	*backoff.ConstantBackOff
}

func NewConstantBackOff(d time.Duration) *constantBackOff {
	origin := backoff.NewConstantBackOff(d)
	eb := &constantBackOff{
		ConstantBackOff: origin,
	}
	eb.backoffDoer = CustomizeBackoff(eb)
	return eb
}
