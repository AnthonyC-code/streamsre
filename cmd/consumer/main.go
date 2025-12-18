// Package main is the entry point for the event consumer.
// The consumer reads review events from Kafka/Redpanda, processes them,
// and writes to Postgres. It exposes /metrics, /healthz, and /readyz endpoints.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"streamsre/internal/config"
	"streamsre/internal/obs"
)

func main() {
	// Initialize logger
	logger, err := obs.NewLogger("info")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	// Load configuration
	cfg, err := config.LoadConsumer()
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	logger.Info("starting consumer",
		zap.Strings("brokers", cfg.KafkaBrokers),
		zap.String("topic", cfg.KafkaTopic),
		zap.String("group_id", cfg.KafkaGroupID),
		zap.Int("metrics_port", cfg.MetricsPort),
	)

	// TODO: Initialize database connection pool
	// TODO: Initialize Kafka consumer
	// TODO: Initialize metrics registry
	// TODO: Start HTTP server for metrics and health checks
	// TODO: Start consumer loop

	// Wait for shutdown signal
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()
	logger.Info("shutting down consumer")

	// TODO: Graceful shutdown (close consumer, db, http server)
}

