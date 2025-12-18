// Package kafka provides dead letter queue (DLQ) functionality.
package kafka

import (
	"context"

	"streamsre/internal/event"
)

// DLQProducer publishes failed events to the dead letter queue topic.
type DLQProducer interface {
	// SendToDLQ publishes a failed event to the DLQ topic.
	SendToDLQ(ctx context.Context, evt *event.DLQEvent) error

	// Close closes the DLQ producer and releases resources.
	Close() error
}

// dlqProducer is the default implementation of DLQProducer.
type dlqProducer struct {
	topic   string
	brokers []string
	// writer *kafkago.Writer // TODO: Add kafka-go writer
}

// NewDLQProducer creates a new DLQ producer.
func NewDLQProducer(brokers []string, mainTopic string) (DLQProducer, error) {
	dlqTopic := mainTopic + DLQTopicSuffix
	_ = dlqTopic // TODO: Use this
	// TODO: Initialize kafka-go writer for DLQ topic
	return nil, nil // TODO
}

// SendToDLQ publishes a failed event to the DLQ topic.
func (p *dlqProducer) SendToDLQ(ctx context.Context, evt *event.DLQEvent) error {
	// TODO: Encode DLQ event and write to DLQ topic
	panic("TODO")
}

// Close closes the DLQ producer and releases resources.
func (p *dlqProducer) Close() error {
	// TODO: Close kafka-go writer
	panic("TODO")
}

