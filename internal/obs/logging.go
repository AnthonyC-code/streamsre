// Package obs provides structured logging using zap.
//
// YOUR TASK (Milestone 2):
// Set up structured JSON logging.
package obs

// TODO: Import "go.uber.org/zap"

// NewLogger creates a production JSON logger.
//
// TODO: Implement:
// 1. Use zap.NewProduction() for JSON output
// 2. Or customize with zap.Config if you want
// func NewLogger() (*zap.Logger, error)

// NewDevelopmentLogger creates a human-readable logger for local dev.
//
// TODO: Implement:
// - Use zap.NewDevelopment()
// func NewDevelopmentLogger() (*zap.Logger, error)

// STRUCTURED LOGGING BEST PRACTICES:
//
// DO include in every log:
//   - event_id (for tracing a specific event)
//   - partition, offset (Kafka position)
//   - latency_ms (how long it took)
//   - result (success, fail, dlq)
//
// Example:
//   logger.Info("event processed",
//       zap.String("event_id", evt.EventID.String()),
//       zap.Int("partition", msg.Partition),
//       zap.Int64("offset", msg.Offset),
//       zap.Duration("latency", time.Since(start)),
//       zap.String("result", "success"),
//   )
//
// DON'T log:
//   - Sensitive data (passwords, tokens)
//   - High-frequency noise (every iteration of a loop)
//   - Entire payloads (can be huge)
