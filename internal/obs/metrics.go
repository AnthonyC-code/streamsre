// Package obs provides observability utilities including metrics and logging.
// Metrics are exposed in Prometheus format for scraping.
package obs

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Metrics holds all application metrics.
type Metrics struct {
	// EventsConsumed counts the total number of events consumed.
	EventsConsumed prometheus.Counter

	// EventsProcessed counts the total number of events successfully processed.
	EventsProcessed prometheus.Counter

	// EventsFailed counts the total number of events that failed processing.
	EventsFailed prometheus.Counter

	// EventsDLQ counts the total number of events sent to the DLQ.
	EventsDLQ prometheus.Counter

	// ProcessingDuration tracks the duration of event processing.
	ProcessingDuration prometheus.Histogram

	// DBQueryDuration tracks the duration of database queries.
	DBQueryDuration *prometheus.HistogramVec

	// KafkaLag tracks the consumer lag per partition.
	KafkaLag *prometheus.GaugeVec

	// RetryCount counts retries per event processing.
	RetryCount prometheus.Counter
}

// NewMetrics creates and registers all application metrics.
func NewMetrics(registry prometheus.Registerer) *Metrics {
	// TODO: Create and register all metrics
	// TODO: Return populated Metrics struct
	panic("TODO")
}

// DefaultMetrics creates metrics with the default Prometheus registry.
func DefaultMetrics() *Metrics {
	return NewMetrics(prometheus.DefaultRegisterer)
}

// RecordEventConsumed increments the events consumed counter.
func (m *Metrics) RecordEventConsumed() {
	// TODO: Increment counter
	panic("TODO")
}

// RecordEventProcessed increments the events processed counter.
func (m *Metrics) RecordEventProcessed() {
	// TODO: Increment counter
	panic("TODO")
}

// RecordEventFailed increments the events failed counter.
func (m *Metrics) RecordEventFailed() {
	// TODO: Increment counter
	panic("TODO")
}

// RecordEventDLQ increments the DLQ counter.
func (m *Metrics) RecordEventDLQ() {
	// TODO: Increment counter
	panic("TODO")
}

// RecordProcessingDuration records the duration of event processing.
func (m *Metrics) RecordProcessingDuration(seconds float64) {
	// TODO: Observe histogram
	panic("TODO")
}

// RecordDBQueryDuration records the duration of a database query.
func (m *Metrics) RecordDBQueryDuration(queryName string, seconds float64) {
	// TODO: Observe histogram with label
	panic("TODO")
}

// SetKafkaLag sets the current lag for a partition.
func (m *Metrics) SetKafkaLag(partition string, lag float64) {
	// TODO: Set gauge with label
	panic("TODO")
}

// RecordRetry increments the retry counter.
func (m *Metrics) RecordRetry() {
	// TODO: Increment counter
	panic("TODO")
}

