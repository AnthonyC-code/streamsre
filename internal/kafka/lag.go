// Package kafka provides consumer lag monitoring.
//
// YOUR TASK (Milestone 10):
// Implement lag measurement for observability.
package kafka

// TODO: Import required packages:
// - "context"

// What is Lag?
// Lag = HighWatermark - CommittedOffset
//
// HighWatermark: Latest message offset in the partition
// CommittedOffset: Where your consumer group has read to
//
// Lag of 0 = You're caught up
// Lag of 1000 = You're 1000 messages behind

// LagInfo contains lag information for a consumer group.
//
// TODO: Define struct with fields:
// - TotalLag      int64            // Sum of all partition lags
// - PartitionLags map[int]int64    // Lag per partition

// LagMonitor measures consumer lag.
//
// TODO: Define struct with fields:
// - brokers []string
// - topic   string
// - groupID string

// NewLagMonitor creates a new lag monitor.
// func NewLagMonitor(brokers []string, topic, groupID string) *LagMonitor

// GetLag returns current lag for the consumer group.
//
// TODO: Implement:
// This is tricky with kafka-go. Options:
//
// Option 1: Use Reader.Stats() (simplest)
//   stats := reader.Stats()
//   return stats.Lag
//
// Option 2: Query manually (more detailed):
//   1. Connect to broker: kafka.Dial("tcp", broker)
//   2. Get partition high watermarks: conn.ReadLastOffset()
//   3. Get consumer group offsets (need to query __consumer_offsets)
//   4. Calculate: lag = highwater - committed
//
// For v1, Option 1 is fine. You can expose reader.Stats().Lag as the metric.
// func (m *LagMonitor) GetLag(ctx context.Context) (*LagInfo, error)

// StartBackground starts a goroutine that samples lag periodically.
//
// TODO: Implement:
// 1. Start a goroutine
// 2. Every 10 seconds, call GetLag()
// 3. Update Prometheus gauge metric
// 4. Stop when context is cancelled
// func (m *LagMonitor) StartBackground(ctx context.Context, metrics *obs.Metrics)
