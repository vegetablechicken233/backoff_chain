# ChainedBackoff

⏱ A fluent chain-style API for the cenkalti/backoff exponential backoff library in Go

## Description

ChainedBackoff provides a fluent chain-style API wrapper around the robust [cenkalti/backoff](https://github.com/cenkalti/backoff) library. This project preserves all the functionality of the original implementation while offering a more ergonomic interface for configuring backoff parameters.

## Overview

This library is a direct modification of cenkalti/backoff, maintaining the same reliable exponential backoff algorithm that multiplicatively adjusts retry intervals. The core improvement is the introduction of a builder pattern that makes configuration more readable and maintainable without changing the underlying implementation.


## Usage

Here’s how to use the **ChainedBackoff** API to configure exponential backoff:

### Original cenkalti/backoff

```go
b := backoff.NewExponentialBackoff()
b.InitialInterval = 100 * time.Millisecond
b.MaxInterval = 10 * time.Second
b.MaxElapsedTime = 1 * time.Minute
receiver , err := backoff.Retry(func() (int, error) {
    // ...
}, opts.WithBackoff(b),opts.WithMaxTries(4))
```

### ChainedBackoff Fluent API

```go
receiver := 0
err := NewExponentialBackOff().
	WithInitialInterval(1*time.Second).
	// ... With other exponential parameters
	WithMaxTries(4).
	// ... With other backoff options
	WithReceiver(&receiver). // receiver set in chain
	Retry(context.Background(), func() (int,error) {
		// ...
	})
fmt.Println(receiver)
```


## License

This project is licensed under the MIT License - see the LICENSE file for details.

