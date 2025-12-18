// Package kafka provides dead letter queue functionality.
//
// YOUR TASK (Milestone 8):
// Implement DLQ for messages that can't be processed.
package kafka

// TODO: Import required packages

// DLQProducer sends failed messages to the dead letter queue.
//
// You have two options for DLQ:
//
// OPTION A: Kafka DLQ Topic
// - Send failed messages to "events.dlq" topic
// - Pro: Easy to replay by consuming DLQ topic
// - Con: Need to manage another topic
//
// OPTION B: Database DLQ Table
// - Insert into dlq_events table
// - Pro: Simpler, can query with SQL
// - Con: Mixing concerns (Kafka issues → DB)
//
// I recommend OPTION A for learning, but either works.

// TODO: Define struct with fields:
// - writer *kafka.Writer  // Writes to DLQ topic

// NewDLQProducer creates a DLQ producer.
//
// TODO: Implement:
// 1. Create kafka.Writer for topic "events.dlq" (or mainTopic + ".dlq")
// 2. Return producer
// func NewDLQProducer(brokers []string, dlqTopic string) *DLQProducer

// SendToDLQ sends a failed message to the DLQ.
//
// TODO: Implement:
// 1. Create a DLQ envelope with:
//    - Original payload (raw bytes)
//    - Reason for failure
//    - Original event ID (if available)
//    - Timestamp
// 2. Encode as JSON
// 3. Write to DLQ topic
// func (p *DLQProducer) SendToDLQ(ctx context.Context, reason string, originalPayload []byte) error

// Close closes the DLQ producer.
// func (p *DLQProducer) Close() error
