// Package service provides retry utilities.
//
// YOUR TASK (Milestone 7):
// Implement exponential backoff with jitter.
package service

// TODO: Import:
// - "context"
// - "math/rand"
// - "time"

// RetryConfig configures retry behavior.
//
// TODO: Define struct with fields:
// - MaxAttempts   int           // Give up after this many tries
// - InitialDelay  time.Duration // First retry delay (e.g., 100ms)
// - MaxDelay      time.Duration // Cap delay at this (e.g., 5s)
// - Multiplier    float64       // Delay multiplier (e.g., 2.0)

// DefaultRetryConfig returns sensible defaults.
// func DefaultRetryConfig() RetryConfig

// Retry executes fn with exponential backoff.
//
// TODO: Implement:
//
// func Retry(ctx context.Context, cfg RetryConfig, fn func() error) error {
//     var lastErr error
//     for attempt := 0; attempt < cfg.MaxAttempts; attempt++ {
//         err := fn()
//         if err == nil {
//             return nil  // Success!
//         }
//         lastErr = err
//
//         // Don't wait after last attempt
//         if attempt == cfg.MaxAttempts-1 {
//             break
//         }
//
//         // Calculate backoff with jitter
//         delay := calculateBackoff(cfg, attempt)
//
//         // Wait (respecting context cancellation)
//         select {
//         case <-ctx.Done():
//             return ctx.Err()
//         case <-time.After(delay):
//         }
//     }
//     return lastErr
// }

// calculateBackoff computes delay for given attempt.
//
// TODO: Implement:
//
// EXPONENTIAL BACKOFF:
// delay = initialDelay * (multiplier ^ attempt)
// attempt 0: 100ms * 2^0 = 100ms
// attempt 1: 100ms * 2^1 = 200ms
// attempt 2: 100ms * 2^2 = 400ms
// ...but capped at MaxDelay
//
// JITTER:
// Without jitter: 1000 requests fail → 1000 retry at same time → fail again
// With jitter: Spread retries randomly to avoid thundering herd
// jitteredDelay = delay * (0.5 + rand.Float64())  // 50-150% of base delay
//
// func calculateBackoff(cfg RetryConfig, attempt int) time.Duration {
//     delay := float64(cfg.InitialDelay) * math.Pow(cfg.Multiplier, float64(attempt))
//     if delay > float64(cfg.MaxDelay) {
//         delay = float64(cfg.MaxDelay)
//     }
//     // Add jitter: 50-150% of calculated delay
//     jitter := 0.5 + rand.Float64()
//     return time.Duration(delay * jitter)
// }
