package backoff_chain

import (
	backoff "github.com/cenkalti/backoff/v5"
)

type zeroBackOff struct {
	*backoffDoer
	*backoff.ZeroBackOff
}

func NewZeroBackoff() *zeroBackOff {
	origin := &backoff.ZeroBackOff{}
	eb := &zeroBackOff{
		ZeroBackOff: origin,
	}
	eb.backoffDoer = CustomizeBackoff(eb)
	return eb
}
