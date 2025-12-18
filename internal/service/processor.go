// Package service contains the business logic for event processing.
// The Processor is responsible for consuming events, processing them,
// and handling failures with retries and DLQ routing.
package service

import (
	"context"

	"go.uber.org/zap"

	"streamsre/internal/db"
	"streamsre/internal/kafka"
	"streamsre/internal/obs"
)

// ProcessorConfig holds configuration for the event processor.
type ProcessorConfig struct {
	// MaxRetries is the maximum number of retries for failed processing.
	MaxRetries int

	// RetryBackoffMs is the initial backoff in milliseconds for retries.
	RetryBackoffMs int
}

// Processor handles event consumption and processing.
type Processor struct {
	config      ProcessorConfig
	consumer    kafka.Consumer
	dlqProducer kafka.DLQProducer
	queries     *db.Queries
	metrics     *obs.Metrics
	logger      *zap.Logger
}

// NewProcessor creates a new event processor.
func NewProcessor(
	config ProcessorConfig,
	consumer kafka.Consumer,
	dlqProducer kafka.DLQProducer,
	queries *db.Queries,
	metrics *obs.Metrics,
	logger *zap.Logger,
) *Processor {
	return &Processor{
		config:      config,
		consumer:    consumer,
		dlqProducer: dlqProducer,
		queries:     queries,
		metrics:     metrics,
		logger:      logger,
	}
}

// Start begins processing events from Kafka.
// This method blocks until the context is cancelled.
func (p *Processor) Start(ctx context.Context) error {
	// TODO: Call consumer.Consume with a handler that:
	// - Checks idempotency (skip if already processed)
	// - Processes the event with retries
	// - Routes to DLQ on permanent failure
	// - Records metrics
	panic("TODO")
}

// processEvent handles a single event with retry logic.
func (p *Processor) processEvent(ctx context.Context, msg *kafka.Message) error {
	// TODO: Check if event already processed (idempotency)
	// TODO: Validate event
	// TODO: Convert to review and insert into database
	// TODO: Mark event as processed
	panic("TODO")
}

// handlePermanentFailure routes a failed event to the DLQ.
func (p *Processor) handlePermanentFailure(ctx context.Context, msg *kafka.Message, reason string) error {
	// TODO: Create DLQ event
	// TODO: Send to DLQ topic
	// TODO: Insert into DLQ table
	panic("TODO")
}

