// Package config provides configuration loading from environment variables.
//
// YOUR TASK (Milestone 2):
// Define config structs and implement loading using github.com/caarlos0/env/v11
package config

// TODO: Import "github.com/caarlos0/env/v11"

// ProducerConfig holds configuration for the event producer.
//
// TODO: Define struct with env tags:
// - Brokers     []string      `env:"KAFKA_BROKERS" envSeparator:"," envDefault:"localhost:9092"`
// - Topic       string        `env:"KAFKA_TOPIC" envDefault:"events.main"`
// - RatePerSec  int           `env:"RATE_PER_SEC" envDefault:"100"`
// - HotKeyPct   float64       `env:"HOT_KEY_PCT" envDefault:"0.8"`
// - Duration    time.Duration `env:"DURATION" envDefault:"60s"`

// ConsumerConfig holds configuration for the event consumer.
//
// TODO: Define struct with env tags:
// - Brokers     []string `env:"KAFKA_BROKERS" envSeparator:"," envDefault:"localhost:9092"`
// - Topic       string   `env:"KAFKA_TOPIC" envDefault:"events.main"`
// - GroupID     string   `env:"KAFKA_GROUP_ID" envDefault:"streamsre-consumer"`
// - MaxInflight int      `env:"MAX_INFLIGHT" envDefault:"64"`
// - DatabaseURL string   `env:"DATABASE_URL" envDefault:"postgres://streamsre:streamsre@localhost:5432/streamsre?sslmode=disable"`
// - MetricsAddr string   `env:"METRICS_ADDR" envDefault:":2112"`

// LoadProducer parses ProducerConfig from environment.
//
// TODO: Implement using env.Parse()
// func LoadProducer() (*ProducerConfig, error)

// LoadConsumer parses ConsumerConfig from environment.
//
// TODO: Implement using env.Parse()
// func LoadConsumer() (*ConsumerConfig, error)
