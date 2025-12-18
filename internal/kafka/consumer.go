// Package kafka provides Kafka consumer functionality.
//
// YOUR TASK (Milestone 5):
// Implement a Kafka consumer with consumer group support.
package kafka

// TODO: Import required packages:
// - "context"
// - "github.com/segmentio/kafka-go"
// - "streamsre/internal/event"

// Message wraps a Kafka message with decoded event.
//
// TODO: Define struct with fields:
// - Event     *event.Event   // Decoded event (nil if decode failed)
// - Raw       []byte         // Original bytes (for DLQ)
// - Partition int
// - Offset    int64
// - Key       []byte

// MessageHandler processes a single message.
// Return nil to commit, return error to retry/DLQ.
type MessageHandler func(ctx context.Context, msg *Message) error

// Consumer reads events from Kafka.
//
// TODO: Define struct with fields:
// - reader *kafka.Reader
// - codec  *event.JSONCodec

// NewConsumer creates a new Kafka consumer with consumer group.
//
// TODO: Implement:
// 1. Create kafka.Reader with:
//    - Brokers: brokers
//    - GroupID: groupID      // THIS ENABLES CONSUMER GROUPS!
//    - Topic: topic
//    - MinBytes: 1
//    - MaxBytes: 10e6        // 10MB
// 2. Return consumer
//
// IMPORTANT: GroupID is what makes this a "consumer group consumer".
// Multiple consumers with same GroupID share partitions.
// Each partition goes to exactly ONE consumer in the group.
// func NewConsumer(brokers []string, topic, groupID string) *Consumer

// Consume reads messages in a loop and calls handler for each.
//
// TODO: Implement:
// 1. Loop forever (until ctx cancelled):
//    - msg, err := reader.ReadMessage(ctx)
//    - Decode message value into Event
//    - Call handler(ctx, &Message{...})
//    - If handler returns nil, commit the message
//    - If handler returns error, DON'T commit (will be redelivered)
//
// COMMIT STRATEGY:
// - Commit AFTER processing, not before!
// - If we crash after processing but before commit, message replays
// - That's okay because we have idempotency in the database
// func (c *Consumer) Consume(ctx context.Context, handler MessageHandler) error

// CommitMessage commits the offset for a message.
//
// TODO: Implement:
// - Call reader.CommitMessages(ctx, kafka.Message{...})
// func (c *Consumer) CommitMessage(ctx context.Context, msg *Message) error

// Close closes the consumer.
//
// TODO: Implement:
// - Call reader.Close()
// func (c *Consumer) Close() error
