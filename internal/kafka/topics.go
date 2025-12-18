// Package kafka provides Kafka/Redpanda client functionality.
// This includes producers, consumers, lag monitoring, and DLQ handling.
package kafka

import "context"

// DefaultTopic is the default topic for review events.
const DefaultTopic = "reviews"

// DLQTopicSuffix is appended to the main topic name for the dead letter queue.
const DLQTopicSuffix = ".dlq"

// TopicConfig holds configuration for creating a topic.
type TopicConfig struct {
	// Name is the topic name.
	Name string

	// NumPartitions is the number of partitions for the topic.
	NumPartitions int

	// ReplicationFactor is the replication factor for the topic.
	ReplicationFactor int

	// RetentionMs is the retention period in milliseconds.
	RetentionMs int64
}

// TopicManager handles topic administration operations.
type TopicManager interface {
	// CreateTopic creates a new topic with the given configuration.
	CreateTopic(ctx context.Context, config TopicConfig) error

	// DeleteTopic deletes the specified topic.
	DeleteTopic(ctx context.Context, name string) error

	// ListTopics returns a list of all topic names.
	ListTopics(ctx context.Context) ([]string, error)

	// TopicExists checks if a topic exists.
	TopicExists(ctx context.Context, name string) (bool, error)
}

// topicManager is the default implementation of TopicManager.
type topicManager struct {
	brokers []string
}

// NewTopicManager creates a new topic manager.
func NewTopicManager(brokers []string) TopicManager {
	return &topicManager{brokers: brokers}
}

// CreateTopic creates a new topic with the given configuration.
func (m *topicManager) CreateTopic(ctx context.Context, config TopicConfig) error {
	// TODO: Use kafka-go to create topic
	panic("TODO")
}

// DeleteTopic deletes the specified topic.
func (m *topicManager) DeleteTopic(ctx context.Context, name string) error {
	// TODO: Use kafka-go to delete topic
	panic("TODO")
}

// ListTopics returns a list of all topic names.
func (m *topicManager) ListTopics(ctx context.Context) ([]string, error) {
	// TODO: Use kafka-go to list topics
	panic("TODO")
}

// TopicExists checks if a topic exists.
func (m *topicManager) TopicExists(ctx context.Context, name string) (bool, error) {
	// TODO: Check if topic exists
	panic("TODO")
}

