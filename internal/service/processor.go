// Package service contains the core business logic.
//
// YOUR TASK (Milestone 5, 6, 7, 8):
// Implement the event processor with:
// - Bounded concurrency (worker pool)
// - Retry with backoff
// - DLQ routing
// - Metrics recording
package service

// TODO: Import required packages

// Processor handles event consumption and processing.
//
// TODO: Define struct with fields:
// - consumer    *kafka.Consumer
// - queries     *db.Queries
// - dlq         *kafka.DLQProducer
// - metrics     *obs.Metrics
// - logger      *zap.Logger
// - maxInflight int
// - maxRetries  int

// NewProcessor creates a new event processor.
// func NewProcessor(...) *Processor

// Start begins processing events.
//
// TODO: Implement the main processing loop:
//
// 1. Create a semaphore for bounded concurrency:
//    sem := make(chan struct{}, p.maxInflight)
//
// 2. Call consumer.Consume(ctx, func(msg) {
//       sem <- struct{}{}  // Acquire (blocks if pool full)
//       go func() {
//           defer func() { <-sem }()  // Release
//           p.processMessage(ctx, msg)
//       }()
//    })
//
// 3. On shutdown (ctx cancelled):
//    - Wait for all in-flight goroutines to finish
//    - Close consumer, DLQ producer
//
// BOUNDED CONCURRENCY EXPLAINED:
// If maxInflight=64 and all workers are busy, the 65th message blocks.
// This is BACKPRESSURE - we slow down to match our processing speed.
// Without this, you'd spawn unlimited goroutines and OOM.
// func (p *Processor) Start(ctx context.Context) error

// processMessage handles a single message with retries.
//
// TODO: Implement:
// 1. Update metrics: inflight++, events_consumed++
// 2. Validate event (if invalid: DLQ immediately, no retry)
// 3. Try processing with retry loop:
//    for attempt := 0; attempt < maxRetries; attempt++ {
//        err := p.processSingle(ctx, msg)
//        if err == nil {
//            success! break
//        }
//        if !isRetryable(err) {
//            DLQ and break
//        }
//        wait with backoff before retry
//        metrics.retries++
//    }
// 4. If all retries exhausted: DLQ
// 5. Update metrics: inflight--, events_processed{result=...}++
// func (p *Processor) processMessage(ctx context.Context, msg *kafka.Message)

// processSingle attempts to process once (no retry logic here).
//
// TODO: Implement:
// 1. Start timer
// 2. Call queries.ProcessEventTx(ctx, eventID, review)
// 3. Record DB latency metric
// 4. Return error (nil on success)
// func (p *Processor) processSingle(ctx context.Context, msg *kafka.Message) error

// isRetryable determines if an error should be retried.
//
// TODO: Implement:
// Retryable errors (transient):
//   - context.DeadlineExceeded (timeout)
//   - Connection refused/reset
//   - "too many connections"
//
// NOT retryable (permanent):
//   - JSON decode error
//   - Validation error
//   - Unique constraint violation
// func isRetryable(err error) bool
