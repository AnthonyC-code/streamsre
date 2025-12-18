// Package kafka provides Kafka producer functionality.
package kafka

import (
	"context"

	"streamsre/internal/event"
)

// ProducerConfig holds configuration for the Kafka producer.
type ProducerConfig struct {
	// Brokers is the list of Kafka broker addresses.
	Brokers []string

	// Topic is the topic to produce messages to.
	Topic string

	// BatchSize is the number of messages to batch before sending.
	BatchSize int

	// BatchTimeoutMs is the maximum time to wait before sending a batch.
	BatchTimeoutMs int
}

// Producer publishes events to Kafka.
type Producer interface {
	// Publish sends a single event to Kafka.
	Publish(ctx context.Context, event *event.ReviewEvent) error

	// PublishBatch sends multiple events to Kafka.
	PublishBatch(ctx context.Context, events []*event.ReviewEvent) error

	// Close closes the producer and releases resources.
	Close() error
}

// producer is the default implementation of Producer.
type producer struct {
	config ProducerConfig
	codec  event.Encoder
	// writer *kafkago.Writer // TODO: Add kafka-go writer
}

// NewProducer creates a new Kafka producer.
func NewProducer(config ProducerConfig, codec event.Encoder) (Producer, error) {
	// TODO: Initialize kafka-go writer with config
	return nil, nil // TODO
}

// Publish sends a single event to Kafka.
func (p *producer) Publish(ctx context.Context, evt *event.ReviewEvent) error {
	// TODO: Encode event and write to Kafka
	panic("TODO")
}

// PublishBatch sends multiple events to Kafka.
func (p *producer) PublishBatch(ctx context.Context, events []*event.ReviewEvent) error {
	// TODO: Encode events and write batch to Kafka
	panic("TODO")
}

// Close closes the producer and releases resources.
func (p *producer) Close() error {
	// TODO: Close kafka-go writer
	panic("TODO")
}

