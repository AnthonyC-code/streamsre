// Package service provides retry utilities with exponential backoff.
package service

import (
	"context"
	"time"
)

// RetryConfig holds configuration for retry behavior.
type RetryConfig struct {
	// MaxRetries is the maximum number of retry attempts.
	MaxRetries int

	// InitialBackoff is the initial backoff duration.
	InitialBackoff time.Duration

	// MaxBackoff is the maximum backoff duration.
	MaxBackoff time.Duration

	// BackoffMultiplier is the multiplier for exponential backoff.
	BackoffMultiplier float64
}

// DefaultRetryConfig returns sensible default retry configuration.
func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxRetries:        3,
		InitialBackoff:    100 * time.Millisecond,
		MaxBackoff:        10 * time.Second,
		BackoffMultiplier: 2.0,
	}
}

// RetryFunc is a function that can be retried.
type RetryFunc func(ctx context.Context) error

// RetryResult contains information about the retry attempt.
type RetryResult struct {
	// Attempts is the number of attempts made.
	Attempts int

	// LastError is the last error encountered.
	LastError error

	// Success indicates if the operation eventually succeeded.
	Success bool
}

// Retry executes the given function with retries according to the config.
func Retry(ctx context.Context, config RetryConfig, fn RetryFunc) RetryResult {
	// TODO: Implement retry logic with exponential backoff
	// - Execute function
	// - On error, wait with backoff and retry
	// - Respect context cancellation
	// - Return result with attempt count
	panic("TODO")
}

// IsRetryable determines if an error should trigger a retry.
func IsRetryable(err error) bool {
	// TODO: Check if error is retryable
	// - Transient errors (network, timeouts) are retryable
	// - Permanent errors (validation, business logic) are not
	panic("TODO")
}

// CalculateBackoff calculates the backoff duration for the given attempt.
func CalculateBackoff(config RetryConfig, attempt int) time.Duration {
	// TODO: Calculate exponential backoff with jitter
	// backoff = min(initialBackoff * (multiplier ^ attempt), maxBackoff)
	panic("TODO")
}

