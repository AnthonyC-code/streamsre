# Runbook: Consumer Lag is Climbing

## Symptoms

- Prometheus alert: `KafkaConsumerLagHigh`
- Grafana dashboard shows increasing lag trend
- `kafka_consumer_lag` metric steadily increasing
- Events taking longer to appear in database after production

## Immediate Checks

1. **Check consumer health**
   ```bash
   # TODO: Check consumer readiness endpoint
   curl http://localhost:2112/readyz
   ```

2. **Check consumer logs for errors**
   ```bash
   # TODO: Tail consumer logs
   make logs | grep consumer
   ```

3. **Check current lag via metrics**
   ```bash
   # TODO: Query Prometheus for current lag
   curl -s 'http://localhost:9090/api/v1/query?query=kafka_consumer_lag' | jq
   ```

4. **Check consumer group status in Redpanda**
   ```bash
   # TODO: Use rpk to check consumer group
   docker exec streamsre-redpanda rpk group describe streamsre-consumer
   ```

5. **Check database health**
   ```bash
   # TODO: Check database connection
   docker exec streamsre-postgres pg_isready -U streamsre
   ```

## Likely Causes

1. **Slow database** - Insert queries taking too long
2. **Consumer crash/restart loop** - Check container status and logs
3. **High message volume** - Producer sending faster than consumer can process
4. **Network issues** - Connectivity between consumer and Kafka/Postgres
5. **Resource exhaustion** - CPU/memory limits on consumer host

## Mitigation

### If database is slow:
```bash
# TODO: Check active queries
docker exec streamsre-postgres psql -U streamsre -c "SELECT * FROM pg_stat_activity WHERE state = 'active';"

# TODO: Check for lock contention
docker exec streamsre-postgres psql -U streamsre -c "SELECT * FROM pg_locks WHERE NOT granted;"
```

### If consumer is overwhelmed:
```bash
# TODO: Scale consumer instances (if using multiple partitions)
# TODO: Increase consumer batch size
# TODO: Check for inefficient processing logic
```

### If producer is too fast:
```bash
# TODO: Consider rate limiting producer
# TODO: Add more consumer instances
# TODO: Increase partition count
```

## Follow-ups

- [ ] Review processing time histograms in Grafana
- [ ] Check if lag is sustained or transient
- [ ] Review database query performance
- [ ] Consider adding consumer scaling automation
- [ ] Update alerting thresholds if necessary

