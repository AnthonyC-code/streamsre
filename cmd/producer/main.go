// Package main is the entry point for the event producer.
// The producer generates review events and publishes them to Kafka/Redpanda.
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
	cfg, err := config.LoadProducer()
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	logger.Info("starting producer",
		zap.Strings("brokers", cfg.KafkaBrokers),
		zap.String("topic", cfg.KafkaTopic),
	)

	// TODO: Initialize Kafka producer
	// TODO: Start event generation loop

	// Wait for shutdown signal
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()
	logger.Info("shutting down producer")

	// TODO: Graceful shutdown
}

