// Package main is the entry point for the event consumer.
//
// YOUR TASK (Milestones 5-11):
// Implement a service that:
// - Consumes from Kafka
// - Processes with bounded concurrency
// - Writes to Postgres with idempotency
// - Exposes metrics and health endpoints
package main

// TODO: Import required packages

// main runs the consumer service.
//
// TODO: Implement:
//
// 1. CONFIGURATION
//    cfg, err := config.LoadConsumer()
//
// 2. LOGGING
//    logger, err := obs.NewLogger()
//    defer logger.Sync()
//
// 3. METRICS
//    metrics := obs.NewMetrics()
//
// 4. DATABASE
//    db, err := db.New(ctx, cfg.DatabaseURL)
//    defer db.Close()
//    queries := db.NewQueries(db)
//
// 5. KAFKA CONSUMER
//    consumer, err := kafka.NewConsumer(cfg.Brokers, cfg.Topic, cfg.GroupID)
//    defer consumer.Close()
//
// 6. DLQ PRODUCER
//    dlq, err := kafka.NewDLQProducer(cfg.Brokers, cfg.Topic+".dlq")
//    defer dlq.Close()
//
// 7. HTTP SERVER (metrics + health)
//    server := service.NewServer(cfg.MetricsAddr, db)
//    server.Start()
//    defer server.Shutdown(ctx)
//
// 8. PROCESSOR
//    proc := service.NewProcessor(...)
//
// 9. SIGNAL HANDLING
//    ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
//    defer cancel()
//
// 10. RUN
//     server.SetReady(true)  // We're ready for traffic
//     logger.Info("consumer started")
//     err := proc.Start(ctx)  // Blocks until shutdown
//
// 11. GRACEFUL SHUTDOWN
//     - Context is cancelled
//     - Processor drains in-flight work
//     - Connections are closed
//     - Exit cleanly

// STARTUP ORDER MATTERS:
// 1. Config first (fail fast if missing)
// 2. Logger next (so we can log errors)
// 3. Database (processor depends on it)
// 4. Kafka (processor depends on it)
// 5. HTTP server (so health checks work)
// 6. Processor (the main work)
//
// SHUTDOWN ORDER IS REVERSE:
// 1. Stop accepting new work (cancel context)
// 2. Drain in-flight work
// 3. Close HTTP server
// 4. Close Kafka
// 5. Close database
