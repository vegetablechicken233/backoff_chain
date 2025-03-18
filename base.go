package backoff_chain

import (
	"backoff"
	"time"

	"github.com/cenkalti/backoff/v5"
)

// Operation is a function that attempts an operation and may be retried.
type Operation func() (any, error)

// Notify is a function called on operation error with the error and backoff duration.
type Notify backoff.Notify

type backoffDoer struct {
	*backoffBase
	b BackOff
}

// backoffBase is the base struct for all backoff implementations.
type backoffBase struct {
	//
	payload *backoffDoer
	//
	notify         Notify        // Optional function to notify on each retry error.
	maxTries       uint          // Maximum number of retry attempts.
	maxElapsedTime time.Duration // Maximum total time for all retries.
	result         any           // result of the last successful operation.
}

// WithNotify sets a notification function to handle retry errors.
func (base *backoffBase) WithNotify(n Notify) *backoffDoer {
	base.notify = n
	return base.payload
}

// WithMaxTries limits the number of all attempts.
func (base *backoffBase) WithMaxTries(n uint) *backoffDoer {
	base.maxTries = n
	return base.payload
}

// WithMaxElapsedTime limits the total duration for retry attempts.
func (base *backoffBase) WithMaxElapsedTime(d time.Duration) *backoffDoer {
	base.maxElapsedTime = d
	return base.payload
}

// WithReceiver save the result of the last successful operation.
// Set the receiver as a pointer to the return value of your operation
// and the result will be saved in the receiver after each successful operation.
//
// Example:
//
//	operation := func() (int, error) {
//		return 1, nil
//	}
//
// var result int
// pl.WithReceiver(&result).Retry(operation)
func (base *backoffBase) WithReceiver(result any) *backoffDoer {
	base.result = result
	return base.payload
}
