package backoff_chain

import (
	backoff "github.com/cenkalti/backoff/v5"
)

type stopBackOff struct {
	*backoffDoer
	*backoff.StopBackOff
}

func NewStopBackOff() *stopBackOff {
	origin := &backoff.StopBackOff{}
	eb := &stopBackOff{
		StopBackOff: origin,
	}
	eb.backoffDoer = CustomizeBackoff(eb)
	return eb
}
