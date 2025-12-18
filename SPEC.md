# StreamSRE Learning Spec

This document is your **implementation guide**. Work through it milestone by milestone. Each milestone builds on the previous one and teaches you core SRE/data-engineering concepts.

**Your goal**: Build a working data pipeline from scratch, understand why each piece exists, and be able to explain it in an interview.

---

## 📚 Concepts You'll Learn

| Concept | Where You'll Apply It |
|---------|----------------------|
| **Message brokers & Kafka** | Producer, Consumer, Topics |
| **Consumer groups & partitions** | Consumer offset management |
| **At-least-once delivery** | Commit strategy + idempotency |
| **Backpressure & concurrency** | Bounded worker pools |
| **Retry patterns** | Exponential backoff with jitter |
| **Dead Letter Queues** | Handling poison messages |
| **Prometheus metrics** | Counters, Gauges, Histograms |
| **Health checks** | Liveness vs Readiness probes |
| **Observability** | Structured logging, dashboards |

---

## 🏗️ Architecture Overview

```
┌──────────────┐     ┌─────────────┐     ┌──────────────┐     ┌──────────────┐
│   Producer   │────▶│  Redpanda   │────▶│   Consumer   │────▶│   Postgres   │
│  (Go CLI)    │     │  (Kafka)    │     │  (Go svc)    │     │              │
└──────────────┘     └─────────────┘     └──────────────┘     └──────────────┘
                           │                    │
                           │              ┌─────┴─────┐
                           │              │  /metrics │
                           │              │  /healthz │
                           │              │  /readyz  │
                           │              └─────┬─────┘
                           │                    │
                     ┌─────┴─────┐        ┌─────┴─────┐
                     │   DLQ     │        │Prometheus │────▶ Grafana
                     │  Topic    │        └───────────┘
                     └───────────┘
```

---

## 🎯 Milestones

### Milestone 0: Understand the Infrastructure (30 min)
**Goal**: Get comfortable with the local environment before writing code.

#### Tasks
- [ ] Run `make up` to start all services
- [ ] Open Grafana at http://localhost:3000 (admin/admin)
- [ ] Open Prometheus at http://localhost:9090
- [ ] Connect to Postgres: `docker exec -it streamsre-postgres psql -U streamsre`
- [ ] Explore Redpanda: `docker exec -it streamsre-redpanda rpk cluster info`

#### Learn: Redpanda/Kafka Basics
Run these commands and understand what they show:

```bash
# List all topics
docker exec streamsre-redpanda rpk topic list

# Create your main topic with 6 partitions
docker exec streamsre-redpanda rpk topic create events.main -p 6

# Create DLQ topic
docker exec streamsre-redpanda rpk topic create events.dlq -p 1

# Describe a topic (shows partitions, replicas)
docker exec streamsre-redpanda rpk topic describe events.main

# Produce a test message manually
echo '{"test": "hello"}' | docker exec -i streamsre-redpanda rpk topic produce events.main

# Consume messages (Ctrl+C to exit)
docker exec streamsre-redpanda rpk topic consume events.main --num 1
```

#### 🧠 Checkpoint Questions
1. What is a partition? Why would you want multiple partitions?
2. What happens when you produce a message without a key vs with a key?
3. What's the difference between a topic and a consumer group?

---

### Milestone 1: Event Model & Codec (1 hour)
**Goal**: Define your data structures and learn JSON serialization.

#### Files to implement
- `internal/event/model.go` - Define the `Event` struct
- `internal/event/codec.go` - Implement JSON encode/decode

#### Event Schema (implement this)
```go
type Event struct {
    EventID       uuid.UUID   `json:"event_id"`       // Unique ID for idempotency
    EventType     string      `json:"event_type"`     // e.g., "review_created"
    Key           string      `json:"key"`            // Partition key, e.g., "user:123"
    Timestamp     time.Time   `json:"ts"`             // When event was created
    SchemaVersion int         `json:"schema_version"` // For future compatibility
    Payload       ReviewData  `json:"payload"`        // The actual data
}

type ReviewData struct {
    ReviewID string `json:"review_id"`
    Rating   int    `json:"rating"`
    Text     string `json:"text"`
}
```

#### Tasks
- [ ] Implement the structs above
- [ ] Write `Encode(event *Event) ([]byte, error)` using `encoding/json`
- [ ] Write `Decode(data []byte) (*Event, error)`
- [ ] Write a simple test in `internal/event/codec_test.go`

#### 🧠 Checkpoint Questions
1. Why do we need `EventID` separate from `ReviewID`?
2. What's the purpose of `SchemaVersion`?
3. Why is `Key` important for Kafka partitioning?

#### Hints
<details>
<summary>Click for hints</summary>

- Use `json.Marshal()` and `json.Unmarshal()`
- For uuid, use `github.com/google/uuid`
- Test with: `go test ./internal/event/...`
</details>

---

### Milestone 2: Configuration (30 min)
**Goal**: Load config from environment variables.

#### Files to implement
- `internal/config/config.go` - Config structs with env tags

#### Tasks
- [ ] Define `ProducerConfig` struct with fields: Brokers, Topic, RatePerSec, HotKeyPct, Duration
- [ ] Define `ConsumerConfig` struct with fields: Brokers, Topic, GroupID, MaxInflight, DatabaseURL, MetricsAddr
- [ ] Use `github.com/caarlos0/env/v11` to parse from environment

#### Example structure
```go
type ConsumerConfig struct {
    Brokers     []string      `env:"KAFKA_BROKERS" envSeparator:"," envDefault:"localhost:9092"`
    Topic       string        `env:"KAFKA_TOPIC" envDefault:"events.main"`
    GroupID     string        `env:"KAFKA_GROUP_ID" envDefault:"streamsre-consumer"`
    MaxInflight int           `env:"MAX_INFLIGHT" envDefault:"64"`
    DatabaseURL string        `env:"DATABASE_URL"`
    MetricsAddr string        `env:"METRICS_ADDR" envDefault:":2112"`
}
```

#### 🧠 Checkpoint Questions
1. Why use environment variables instead of config files for production?
2. What's a reasonable default for `MaxInflight`? Why not 1000?

---

### Milestone 3: Database Layer (1 hour)
**Goal**: Connect to Postgres and implement the queries you need.

#### Files to implement
- `internal/db/db.go` - Connection pool setup
- `internal/db/queries.go` - SQL operations

#### Tables (already in migrations/001_init.sql)
```sql
-- Idempotency: Have we processed this event before?
processed_events(event_id UUID PRIMARY KEY, processed_at TIMESTAMPTZ)

-- Sink: Where processed data lands
reviews(review_id TEXT PRIMARY KEY, user_key TEXT, rating INT, text TEXT, created_at TIMESTAMPTZ)

-- DLQ: Failed events for later analysis
dlq_events(event_id UUID PRIMARY KEY, reason TEXT, original_payload JSONB, failed_at TIMESTAMPTZ)
```

#### Tasks
- [ ] Create connection pool using `pgxpool.New(ctx, databaseURL)`
- [ ] Implement `IsEventProcessed(ctx, eventID) (bool, error)`
- [ ] Implement `MarkEventProcessed(ctx, eventID) error`
- [ ] Implement `InsertReview(ctx, review) error`
- [ ] Implement `InsertDLQEvent(ctx, eventID, reason, payload) error`
- [ ] Implement `ProcessEventTx(ctx, eventID, review) error` - **CRITICAL**: This must be transactional!

#### The Critical Transaction Pattern
```go
// ProcessEventTx must do these steps atomically (in one transaction):
// 1. Check if event already processed (SELECT FROM processed_events)
// 2. If already processed, return early (skip - idempotent!)
// 3. Insert the review
// 4. Mark event as processed
// 5. Commit transaction
//
// If any step fails, rollback everything.
// This is how you achieve "exactly-once" semantics with "at-least-once" delivery.
```

#### 🧠 Checkpoint Questions
1. Why do we need a transaction here? What could go wrong without one?
2. What happens if the consumer crashes after writing to `reviews` but before committing?
3. How does `processed_events` prevent duplicate reviews?

#### Hints
<details>
<summary>Click for hints</summary>

```go
// Start a transaction
tx, err := pool.Begin(ctx)
if err != nil { return err }
defer tx.Rollback(ctx) // Rollback if we don't commit

// Do your queries using tx.Exec() or tx.QueryRow()

// Commit at the end
return tx.Commit(ctx)
```
</details>

---

### Milestone 4: Kafka Producer (1 hour)
**Goal**: Produce messages to Kafka with controlled rate and key distribution.

#### Files to implement
- `internal/kafka/producer.go` - Kafka writer wrapper
- `cmd/producer/main.go` - CLI that produces events

#### Tasks
- [ ] Create a `kafka.Writer` with your config
- [ ] Implement `Produce(ctx, event) error` that encodes and writes
- [ ] In main.go, implement a loop that:
  - Generates random events at configured rate (use `time.Ticker`)
  - Uses "hot keys" for some percentage (simulates real traffic patterns)
  - Runs for configured duration then exits

#### Hot Key Simulation
Real systems have "hot partitions" - some users are way more active:
```go
// 80% of traffic goes to 20% of keys (configurable via HOT_KEY_PCT)
hotKeys := []string{"user:1", "user:2", "user:3"}  // Hot users
coldKeys := []string{"user:100", "user:101", ...}  // Everyone else

if rand.Float64() < hotKeyPct {
    key = hotKeys[rand.Intn(len(hotKeys))]
} else {
    key = coldKeys[rand.Intn(len(coldKeys))]
}
```

#### 🧠 Checkpoint Questions
1. What happens when all messages have the same key?
2. How does Kafka decide which partition gets a message?
3. Why would hot keys cause problems in production?

#### Test It
```bash
# Run producer
KAFKA_TOPIC=events.main RATE_PER_SEC=10 DURATION=30s go run ./cmd/producer

# In another terminal, watch messages arrive
docker exec streamsre-redpanda rpk topic consume events.main
```

---

### Milestone 5: Kafka Consumer Basics (1.5 hours)
**Goal**: Read messages from Kafka with a consumer group.

#### Files to implement
- `internal/kafka/consumer.go` - Kafka reader wrapper
- `cmd/consumer/main.go` - Service that consumes events

#### Tasks
- [ ] Create a `kafka.Reader` with GroupID (this enables consumer groups!)
- [ ] Implement `Consume(ctx, handler func(Message) error) error`
- [ ] In main.go, start consuming and just log each message for now

#### Key Concepts to Understand

**Consumer Groups**: Multiple consumers with same GroupID share the work. Each partition is assigned to exactly one consumer in the group.

**Offsets**: Kafka tracks where each consumer group has read to. When you "commit", you're saying "I've processed up to here."

**Commit Strategy** (implement this):
```go
// WRONG: Commit before processing
reader.CommitMessages(ctx, msg)  // <- Danger!
processMessage(msg)              // If this fails, message is lost

// RIGHT: Commit after processing
processMessage(msg)              // Process first
reader.CommitMessages(ctx, msg)  // Then commit

// If we crash between process and commit, message replays.
// That's why we need idempotency!
```

#### 🧠 Checkpoint Questions
1. What happens if two consumers have different GroupIDs and read the same topic?
2. What happens if a consumer crashes without committing?
3. Why do we need idempotency if we commit after processing?

#### Test It
```bash
# Terminal 1: Run producer
go run ./cmd/producer

# Terminal 2: Run consumer
go run ./cmd/consumer

# You should see consumer logging each message
```

---

### Milestone 6: Bounded Concurrency (1.5 hours)
**Goal**: Process messages concurrently with backpressure.

#### The Problem
If you process messages one at a time, you're slow. If you spawn unlimited goroutines, you'll OOM or overwhelm the database.

#### Solution: Worker Pool Pattern
```go
// Bounded worker pool with semaphore
sem := make(chan struct{}, maxInflight)  // Limit concurrent workers

for msg := range messages {
    sem <- struct{}{}  // Acquire slot (blocks if pool is full = backpressure!)
    
    go func(m Message) {
        defer func() { <-sem }()  // Release slot when done
        processMessage(m)
    }(msg)
}
```

#### Tasks
- [ ] Implement worker pool in your consumer
- [ ] Track inflight count as a metric (we'll wire metrics later)
- [ ] Handle graceful shutdown (wait for inflight to drain)

#### 🧠 Checkpoint Questions
1. What happens if database is slow and all 64 workers are busy?
2. How does this prevent the consumer from falling behind Kafka?
3. What's the tradeoff of setting MaxInflight too high vs too low?

---

### Milestone 7: Retry with Backoff (1 hour)
**Goal**: Handle transient failures gracefully.

#### Files to implement
- `internal/service/retry.go` - Retry logic with exponential backoff + jitter

#### The Pattern
```
Attempt 1: immediate
Attempt 2: wait 100ms + jitter
Attempt 3: wait 200ms + jitter
Attempt 4: wait 400ms + jitter
...
Attempt 8: give up → send to DLQ
```

#### Why Jitter?
Without jitter, if 1000 requests fail at the same time, they all retry at the same time → thundering herd → fail again → repeat.

Jitter spreads out retries randomly.

#### Tasks
- [ ] Implement `Retry(ctx, maxAttempts, fn func() error) error`
- [ ] Use exponential backoff: `backoff = min(baseDelay * 2^attempt, maxDelay)`
- [ ] Add jitter: `actualDelay = backoff * (0.5 + rand.Float64())`
- [ ] Respect context cancellation (don't retry if shutting down)

#### 🧠 Checkpoint Questions
1. Which errors should you retry? Which shouldn't you?
2. Why cap the backoff at a maximum (e.g., 5 seconds)?
3. What happens if you retry forever instead of giving up after N attempts?

#### Hints
<details>
<summary>Distinguishing Retryable Errors</summary>

```go
// Retryable: transient failures
// - Network timeout
// - Connection refused (database might restart)
// - "too many connections"

// NOT retryable: permanent failures
// - Invalid JSON (will never parse correctly)
// - Constraint violation (duplicate key)
// - 400 Bad Request
```
</details>

---

### Milestone 8: Dead Letter Queue (1 hour)
**Goal**: Handle poison messages that can never be processed.

#### What Goes to DLQ?
1. **Invalid payload**: Can't decode JSON, missing required fields
2. **Max retries exceeded**: Tried N times, still failing
3. **Business rule violation**: Rating is -5 (impossible value)

#### Tasks
- [ ] Implement `SendToDLQ(ctx, eventID, reason, originalPayload) error`
- [ ] Option A: Write to Kafka topic `events.dlq`
- [ ] Option B: Write to Postgres table `dlq_events`
- [ ] Wire into your processor: invalid → DLQ immediately, max retries → DLQ

#### 🧠 Checkpoint Questions
1. Why not just drop poison messages?
2. How would you "replay" messages from the DLQ after fixing a bug?
3. Should DLQ messages be retried automatically? Why or why not?

---

### Milestone 9: Metrics (2 hours)
**Goal**: Expose Prometheus metrics for observability.

#### Files to implement
- `internal/obs/metrics.go` - Define and register all metrics

#### Metrics to Implement

| Metric | Type | Labels | Purpose |
|--------|------|--------|---------|
| `events_consumed_total` | Counter | - | How many messages pulled from Kafka |
| `events_processed_total` | Counter | `result` | success/fail/dlq outcomes |
| `processing_duration_seconds` | Histogram | - | How long each event takes |
| `db_query_duration_seconds` | Histogram | `query` | Database latency |
| `retries_total` | Counter | - | How many retries happened |
| `inflight_workers` | Gauge | - | Current active goroutines |
| `consumer_lag` | Gauge | `partition` | Messages behind |

#### Tasks
- [ ] Create each metric using `prometheus.NewCounter()`, etc.
- [ ] Register with `prometheus.MustRegister()`
- [ ] Instrument your code to record metrics at the right places
- [ ] Expose `/metrics` endpoint using `promhttp.Handler()`

#### Example: Timing a function
```go
start := time.Now()
err := doSomething()
duration := time.Since(start).Seconds()
metrics.ProcessingDuration.Observe(duration)
```

#### 🧠 Checkpoint Questions
1. What's the difference between Counter, Gauge, and Histogram?
2. Why use a Histogram for latency instead of just an average?
3. What labels would be useful vs harmful? (Hint: cardinality)

#### Test It
```bash
# Run consumer
go run ./cmd/consumer

# Check metrics
curl localhost:2112/metrics | grep streamsre
```

---

### Milestone 10: Consumer Lag (1 hour)
**Goal**: Measure how far behind the consumer is.

#### What is Lag?
```
Lag = Latest Offset in Partition - Consumer's Committed Offset

If partition has 1000 messages and you've processed 800:
Lag = 1000 - 800 = 200 messages behind
```

#### Tasks
- [ ] Implement lag measurement in `internal/kafka/lag.go`
- [ ] Run a background goroutine that samples lag every 10 seconds
- [ ] Expose as `consumer_lag` gauge metric

#### How to Get Lag with kafka-go
```go
// Option 1: Use Reader Stats (simpler)
stats := reader.Stats()
// stats.Lag gives you total lag

// Option 2: Query offsets directly (more detailed)
// Get partition high watermarks via Conn
// Get consumer group committed offsets
// Calculate difference
```

#### 🧠 Checkpoint Questions
1. Is lag of 1000 always bad? When might it be okay?
2. What could cause lag to grow unboundedly?
3. How would you alert on lag?

---

### Milestone 11: Health Endpoints (30 min)
**Goal**: Implement Kubernetes-style health checks.

#### Files to implement
- `internal/service/http.go` - HTTP server with health endpoints

#### Endpoints

**`/healthz` (Liveness)**
- Returns 200 if the process is running
- Kubernetes uses this to know if it should restart the container
- Should NOT check dependencies (database, Kafka)

**`/readyz` (Readiness)**
- Returns 200 if ready to receive traffic
- Check: Can I reach Kafka? Can I reach Postgres?
- If unhealthy, Kubernetes removes from load balancer (no traffic)

#### Tasks
- [ ] Implement `/healthz` - just return `{"status": "ok"}`
- [ ] Implement `/readyz` - ping Kafka and Postgres, return 200 or 503
- [ ] Implement `/metrics` - serve Prometheus metrics

#### 🧠 Checkpoint Questions
1. Why separate liveness from readiness?
2. What happens if `/readyz` always returns 500?
3. Should health checks have timeouts?

---

### Milestone 12: Grafana Dashboard (1 hour)
**Goal**: Visualize your metrics.

#### Tasks
- [ ] Open Grafana at http://localhost:3000
- [ ] Create a new dashboard
- [ ] Add panels for:
  - Consumer lag (timeseries)
  - Events processed per second (rate)
  - Processing latency p50/p95/p99 (histogram_quantile)
  - Error rate (events failed / events total)
  - Inflight workers (gauge)
  - DB latency heatmap

#### PromQL Examples
```promql
# Rate of events processed per second
rate(events_processed_total[1m])

# 95th percentile processing latency
histogram_quantile(0.95, rate(processing_duration_seconds_bucket[5m]))

# Consumer lag
consumer_lag

# Error rate
rate(events_processed_total{result="fail"}[5m]) / rate(events_processed_total[5m])
```

#### 🧠 Checkpoint Questions
1. Why use `rate()` instead of just the counter value?
2. What does `[1m]` mean in PromQL?
3. How would you set up an alert for "lag > 10000 for 5 minutes"?

---

### Milestone 13: Chaos Engineering (1.5 hours)
**Goal**: Break things on purpose and verify your system handles it.

#### Drill 1: Lag Climbing
```bash
# Produce faster than consumer can handle
RATE_PER_SEC=1000 go run ./cmd/producer

# Watch lag climb in Grafana
# Then reduce rate or increase MAX_INFLIGHT
```

**Document in runbooks/lag_is_climbing.md**:
- What metrics spiked?
- How did you diagnose the root cause?
- What was the mitigation?

#### Drill 2: Slow Database
```bash
# Add artificial latency to Postgres
docker exec streamsre-postgres psql -U streamsre -c "
  CREATE OR REPLACE FUNCTION slow_trigger() RETURNS trigger AS \$\$
  BEGIN
    PERFORM pg_sleep(2);
    RETURN NEW;
  END;
  \$\$ LANGUAGE plpgsql;

  CREATE TRIGGER slow_insert BEFORE INSERT ON reviews
  FOR EACH ROW EXECUTE FUNCTION slow_trigger();
"

# Run your pipeline and watch retries spike
# Clean up:
docker exec streamsre-postgres psql -U streamsre -c "DROP TRIGGER slow_insert ON reviews;"
```

#### Drill 3: Poison Message
```bash
# Produce invalid JSON
echo 'not valid json at all' | docker exec -i streamsre-redpanda rpk topic produce events.main

# Verify it goes to DLQ
# Check dlq_events table or events.dlq topic
```

#### Drill 4: Consumer Crash
```bash
# Start consumer, then kill it mid-processing
go run ./cmd/consumer &
PID=$!
sleep 5
kill -9 $PID

# Restart consumer - verify no duplicates in reviews table
# (idempotency should prevent double-writes)
```

---

### Milestone 14: Write Runbooks & Postmortems (1 hour)
**Goal**: Document what you learned from the drills.

#### Runbook Template (already in runbooks/)
For each failure mode, document:
1. **Symptoms**: What alerts fire? What does Grafana show?
2. **Immediate Checks**: Commands to run, queries to execute
3. **Likely Causes**: What usually causes this?
4. **Mitigation**: How to fix it RIGHT NOW
5. **Follow-ups**: What to do after the incident

#### Postmortem Template
Create `postmortems/YYYY-MM-DD-title.md`:
```markdown
# Postmortem: [Title]

## Summary
One paragraph: what happened, impact, duration.

## Timeline
- HH:MM - First alert
- HH:MM - Investigation started
- HH:MM - Root cause identified
- HH:MM - Mitigation applied
- HH:MM - Resolved

## Root Cause
Why did this happen?

## Impact
What was affected? How many events delayed/lost?

## What Went Well
- Alerts fired correctly
- Runbook was helpful

## What Could Be Improved
- Need better monitoring for X
- Should add circuit breaker

## Action Items
- [ ] Add alert for Y
- [ ] Implement Z
```

---

## ✅ Definition of Done

You're finished when:

- [ ] `docker compose up` starts all infrastructure
- [ ] Producer generates events at configurable rate
- [ ] Consumer processes events and writes to Postgres
- [ ] Killing consumer doesn't cause duplicate reviews (idempotency works!)
- [ ] Invalid messages go to DLQ
- [ ] `/metrics` exposes lag, latency, throughput, errors
- [ ] Grafana dashboard shows key metrics
- [ ] 3 runbooks completed with real examples
- [ ] 2 postmortems from chaos drills

---

## 🚀 Stretch Goals (if you want more)

1. **Circuit Breaker**: Stop calling Postgres if it's consistently failing
2. **Rate Limiting**: Consumer limits its own throughput
3. **Schema Registry**: Validate events against a schema
4. **Multi-Consumer**: Run 2+ consumers and watch partition rebalancing
5. **Exactly-Once**: Use Kafka transactions (advanced!)
6. **Kubernetes**: Deploy to minikube with proper manifests

---

## 📖 Resources

### Kafka
- [Kafka: The Definitive Guide](https://www.confluent.io/resources/kafka-the-definitive-guide/) (free ebook)
- [kafka-go documentation](https://pkg.go.dev/github.com/segmentio/kafka-go)

### Prometheus
- [Prometheus docs](https://prometheus.io/docs/introduction/overview/)
- [PromQL basics](https://prometheus.io/docs/prometheus/latest/querying/basics/)

### Go Patterns
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Concurrency Patterns](https://go.dev/blog/pipelines)

### SRE
- [Google SRE Book](https://sre.google/sre-book/table-of-contents/) (free online)
- [The Site Reliability Workbook](https://sre.google/workbook/table-of-contents/)

---

Good luck! Start with Milestone 0 to get comfortable with the environment, then work through each milestone in order. Don't skip ahead - each one builds on the previous.

