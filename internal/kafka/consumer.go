// Package kafka provides Kafka consumer functionality.
package kafka

import (
	"context"

	"streamsre/internal/event"
)

// ConsumerConfig holds configuration for the Kafka consumer.
type ConsumerConfig struct {
	// Brokers is the list of Kafka broker addresses.
	Brokers []string

	// Topic is the topic to consume messages from.
	Topic string

	// GroupID is the consumer group ID.
	GroupID string

	// MinBytes is the minimum number of bytes to fetch per request.
	MinBytes int

	// MaxBytes is the maximum number of bytes to fetch per request.
	MaxBytes int

	// CommitInterval is the interval at which to commit offsets.
	CommitIntervalMs int
}

// Message represents a consumed Kafka message.
type Message struct {
	// Event is the decoded review event.
	Event *event.ReviewEvent

	// Partition is the partition the message was read from.
	Partition int

	// Offset is the offset of the message in the partition.
	Offset int64

	// Key is the message key.
	Key []byte

	// Raw is the raw message bytes.
	Raw []byte
}

// MessageHandler is a function that processes a consumed message.
type MessageHandler func(ctx context.Context, msg *Message) error

// Consumer reads events from Kafka.
type Consumer interface {
	// Consume starts consuming messages and calls the handler for each message.
	// This method blocks until the context is cancelled.
	Consume(ctx context.Context, handler MessageHandler) error

	// CommitMessage commits the offset for the given message.
	CommitMessage(ctx context.Context, msg *Message) error

	// Close closes the consumer and releases resources.
	Close() error
}

// consumer is the default implementation of Consumer.
type consumer struct {
	config ConsumerConfig
	codec  event.Decoder
	// reader *kafkago.Reader // TODO: Add kafka-go reader
}

// NewConsumer creates a new Kafka consumer.
func NewConsumer(config ConsumerConfig, codec event.Decoder) (Consumer, error) {
	// TODO: Initialize kafka-go reader with config
	return nil, nil // TODO
}

// Consume starts consuming messages and calls the handler for each message.
func (c *consumer) Consume(ctx context.Context, handler MessageHandler) error {
	// TODO: Read messages from Kafka in a loop
	// - Decode each message
	// - Call handler
	// - Handle errors and commit offsets
	panic("TODO")
}

// CommitMessage commits the offset for the given message.
func (c *consumer) CommitMessage(ctx context.Context, msg *Message) error {
	// TODO: Commit offset for message
	panic("TODO")
}

// Close closes the consumer and releases resources.
func (c *consumer) Close() error {
	// TODO: Close kafka-go reader
	panic("TODO")
}

