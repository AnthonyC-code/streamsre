// Package obs provides observability: metrics and logging.
//
// YOUR TASK (Milestone 9):
// Implement Prometheus metrics for the consumer.
package obs

// TODO: Import required packages:
// - "github.com/prometheus/client_golang/prometheus"

// Metrics holds all Prometheus metrics for the application.
//
// TODO: Define struct with these metrics:
//
// COUNTERS (things that only go up):
// - EventsConsumed     prometheus.Counter      // Messages pulled from Kafka
// - EventsProcessed    *prometheus.CounterVec  // Labeled by result: success, fail, dlq
// - RetriesTotal       prometheus.Counter      // Number of retry attempts
//
// HISTOGRAMS (distributions - latency, etc.):
// - ProcessingDuration prometheus.Histogram    // How long each event takes
// - DBQueryDuration    *prometheus.HistogramVec // DB latency by query type
//
// GAUGES (values that go up and down):
// - InflightWorkers    prometheus.Gauge        // Current active goroutines
// - ConsumerLag        *prometheus.GaugeVec    // Lag per partition

// NewMetrics creates and registers all metrics.
//
// TODO: Implement:
// 1. Create each metric using prometheus.New*()
// 2. Register with prometheus.MustRegister()
// 3. Return &Metrics{...}
//
// Example:
//   eventsConsumed := prometheus.NewCounter(prometheus.CounterOpts{
//       Name: "streamsre_events_consumed_total",
//       Help: "Total number of events consumed from Kafka",
//   })
//   prometheus.MustRegister(eventsConsumed)
//
// For labeled metrics (CounterVec, GaugeVec):
//   eventsProcessed := prometheus.NewCounterVec(prometheus.CounterOpts{
//       Name: "streamsre_events_processed_total",
//       Help: "Total events processed, by result",
//   }, []string{"result"})  // Labels
//   prometheus.MustRegister(eventsProcessed)
//
// func NewMetrics() *Metrics

// Recording methods:
//
// TODO: Implement helper methods like:
// - m.RecordEventConsumed()
// - m.RecordEventProcessed(result string)  // "success", "fail", "dlq"
// - m.RecordProcessingDuration(seconds float64)
// - m.SetInflightWorkers(count int)
// - m.SetConsumerLag(partition int, lag int64)

// PROMETHEUS METRIC TYPES EXPLAINED:
//
// Counter: Only goes up. Reset when process restarts.
//   Use for: requests, events, errors
//   Query with: rate(metric_name[5m]) to get per-second rate
//
// Gauge: Goes up and down.
//   Use for: current queue size, active connections, temperature
//   Query with: metric_name (current value)
//
// Histogram: Distribution of values, pre-aggregated into buckets.
//   Use for: latency, request size
//   Query with: histogram_quantile(0.95, rate(metric_name_bucket[5m]))
//
// Labels: Add dimensions to metrics.
//   Good: {result="success"}, {method="GET"}, {status="200"}
//   Bad: {user_id="..."} - Too many unique values (high cardinality)!
