// Package config provides configuration loading for producer and consumer.
// Configuration is loaded from environment variables using the env library.
package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

// Producer holds configuration for the event producer.
type Producer struct {
	// KafkaBrokers is a list of Kafka broker addresses.
	KafkaBrokers []string `env:"KAFKA_BROKERS" envSeparator:"," envDefault:"localhost:9092"`

	// KafkaTopic is the topic to produce events to.
	KafkaTopic string `env:"KAFKA_TOPIC" envDefault:"reviews"`

	// LogLevel sets the logging verbosity (debug, info, warn, error).
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`

	// EventIntervalMs is the interval between generated events in milliseconds.
	EventIntervalMs int `env:"EVENT_INTERVAL_MS" envDefault:"1000"`
}

// Consumer holds configuration for the event consumer.
type Consumer struct {
	// KafkaBrokers is a list of Kafka broker addresses.
	KafkaBrokers []string `env:"KAFKA_BROKERS" envSeparator:"," envDefault:"localhost:9092"`

	// KafkaTopic is the topic to consume events from.
	KafkaTopic string `env:"KAFKA_TOPIC" envDefault:"reviews"`

	// KafkaGroupID is the consumer group ID.
	KafkaGroupID string `env:"KAFKA_GROUP_ID" envDefault:"streamsre-consumer"`

	// DatabaseURL is the Postgres connection string.
	DatabaseURL string `env:"DATABASE_URL" envDefault:"postgres://streamsre:streamsre@localhost:5432/streamsre?sslmode=disable"`

	// MetricsPort is the port for the metrics and health HTTP server.
	MetricsPort int `env:"METRICS_PORT" envDefault:"2112"`

	// LogLevel sets the logging verbosity (debug, info, warn, error).
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`

	// MaxRetries is the maximum number of retries for failed event processing.
	MaxRetries int `env:"MAX_RETRIES" envDefault:"3"`

	// RetryBackoffMs is the initial backoff duration in milliseconds for retries.
	RetryBackoffMs int `env:"RETRY_BACKOFF_MS" envDefault:"100"`
}

// LoadProducer loads producer configuration from environment variables.
func LoadProducer() (*Producer, error) {
	cfg := &Producer{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parsing producer config: %w", err)
	}
	return cfg, nil
}

// LoadConsumer loads consumer configuration from environment variables.
func LoadConsumer() (*Consumer, error) {
	cfg := &Consumer{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parsing consumer config: %w", err)
	}
	return cfg, nil
}

