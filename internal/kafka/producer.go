// Package kafka provides Kafka producer functionality.
//
// YOUR TASK (Milestone 4):
// Implement a Kafka producer using github.com/segmentio/kafka-go
package kafka

// TODO: Import required packages:
// - "context"
// - "github.com/segmentio/kafka-go"
// - "streamsre/internal/event"

// Producer writes events to Kafka.
//
// TODO: Define struct with fields:
// - writer *kafka.Writer
// - codec  *event.JSONCodec  // Or however you named your codec

// NewProducer creates a new Kafka producer.
//
// TODO: Implement:
// 1. Create kafka.Writer with:
//    - Addr: kafka.TCP(brokers...)
//    - Topic: topic
//    - Balancer: &kafka.Hash{}  // Partition by key
// 2. Return producer
// func NewProducer(brokers []string, topic string) *Producer

// Produce sends an event to Kafka.
//
// TODO: Implement:
// 1. Encode event to JSON using codec
// 2. Create kafka.Message{Key: []byte(event.Key), Value: jsonBytes}
// 3. Call writer.WriteMessages(ctx, message)
// func (p *Producer) Produce(ctx context.Context, evt *event.Event) error

// Close closes the producer.
//
// TODO: Implement:
// - Call p.writer.Close()
// func (p *Producer) Close() error
