# StreamSRE

A local data-infrastructure SRE lab project demonstrating event streaming with Go, Kafka (Redpanda), and Postgres.

## Architecture

```
Go Producer → Redpanda (Kafka) → Go Consumer → Postgres
                                      ↓
                              Prometheus ← Grafana
```

## Prerequisites

- Go 1.22+
- Docker & Docker Compose
- Make

## Quick Start

### 1. Start Infrastructure

```bash
make up
```

This starts:
- **Redpanda** (Kafka API): `localhost:9092`
- **Postgres**: `localhost:5432`
- **Prometheus**: `localhost:9090`
- **Grafana**: `localhost:3000` (admin/admin)

### 2. Run Migrations

```bash
make migrate
```

### 3. Run the Consumer

```bash
make run-consumer
```

### 4. Run the Producer

```bash
make run-producer
```

## Environment Variables

### Producer

| Variable | Description | Default |
|----------|-------------|---------|
| `KAFKA_BROKERS` | Comma-separated Kafka broker addresses | `localhost:9092` |
| `KAFKA_TOPIC` | Topic to produce events to | `reviews` |
| `LOG_LEVEL` | Logging level (debug, info, warn, error) | `info` |

### Consumer

| Variable | Description | Default |
|----------|-------------|---------|
| `KAFKA_BROKERS` | Comma-separated Kafka broker addresses | `localhost:9092` |
| `KAFKA_TOPIC` | Topic to consume events from | `reviews` |
| `KAFKA_GROUP_ID` | Consumer group ID | `streamsre-consumer` |
| `DATABASE_URL` | Postgres connection string | `postgres://streamsre:streamsre@localhost:5432/streamsre?sslmode=disable` |
| `METRICS_PORT` | Port for metrics/health endpoints | `2112` |
| `LOG_LEVEL` | Logging level | `info` |

## Endpoints (Consumer)

- `GET /metrics` - Prometheus metrics
- `GET /healthz` - Liveness probe
- `GET /readyz` - Readiness probe

## Project Structure

```
streamsre/
├── cmd/
│   ├── producer/     # Event producer CLI
│   └── consumer/     # Event consumer service
├── internal/
│   ├── config/       # Configuration loading
│   ├── event/        # Event models and codecs
│   ├── kafka/        # Kafka producer/consumer
│   ├── db/           # Database access layer
│   ├── obs/          # Observability (metrics, logging)
│   └── service/      # Business logic and HTTP handlers
├── configs/          # Prometheus, Grafana configs
├── migrations/       # SQL migrations
├── runbooks/         # Incident response guides
└── scripts/          # Helper and chaos scripts
```

## Runbooks

- [Lag is Climbing](runbooks/lag_is_climbing.md)
- [DB Slow/Retries](runbooks/db_slow_retries.md)
- [Poison Messages in DLQ](runbooks/poison_messages_dlq.md)

## Next Steps

- [ ] Implement producer event generation logic
- [ ] Implement consumer message processing
- [ ] Add idempotency checks via `processed_events` table
- [ ] Implement DLQ routing for poison messages
- [ ] Create Grafana dashboard panels
- [ ] Add alerting rules to Prometheus
- [ ] Implement graceful shutdown
- [ ] Add integration tests
- [ ] Set up CI/CD pipeline

## License

MIT

