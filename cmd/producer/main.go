// Package main is the entry point for the event producer.
//
// YOUR TASK (Milestone 4):
// Implement a CLI that produces events at a configurable rate.
package main

// TODO: Import required packages

// main produces events to Kafka.
//
// TODO: Implement:
//
// 1. Load config using config.LoadProducer()
//
// 2. Set up logger using obs.NewLogger()
//
// 3. Create Kafka producer using kafka.NewProducer()
//
// 4. Set up signal handling for graceful shutdown:
//    ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
//    defer cancel()
//
// 5. Create a ticker for rate limiting:
//    interval := time.Second / time.Duration(cfg.RatePerSec)
//    ticker := time.NewTicker(interval)
//
// 6. Run the main loop:
//    timeout := time.After(cfg.Duration)
//    for {
//        select {
//        case <-ctx.Done():
//            return  // Shutdown
//        case <-timeout:
//            return  // Duration elapsed
//        case <-ticker.C:
//            event := generateEvent(cfg.HotKeyPct)
//            producer.Produce(ctx, event)
//        }
//    }
//
// 7. Close producer on exit

// generateEvent creates a random event.
//
// TODO: Implement:
// 1. Generate random event ID using uuid.New()
// 2. Pick a key based on hot key percentage:
//    - hotKeyPct chance: pick from hot keys ("user:1", "user:2", etc.)
//    - else: pick from cold keys ("user:100", "user:101", etc.)
// 3. Generate random rating (1-5), random text
// 4. Return the event
//
// WHY HOT KEYS?
// In real systems, some users are much more active than others.
// This creates "hot partitions" - some partitions get way more traffic.
// It's important to understand because it affects scaling.
// func generateEvent(hotKeyPct float64) *event.Event
