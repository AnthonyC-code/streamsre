# Runbook: Database Slow / High Retry Rate

## Symptoms

- Prometheus alert: `DBQueryLatencyHigh` or `RetryRateHigh`
- Grafana shows elevated `db_query_duration_seconds` histogram
- `retry_count` metric increasing
- Consumer logs show database timeout errors
- Consumer lag may be climbing as a secondary effect

## Immediate Checks

1. **Check database connectivity**
   ```bash
   # TODO: Check Postgres is responding
   docker exec streamsre-postgres pg_isready -U streamsre
   ```

2. **Check active connections**
   ```bash
   # TODO: Check connection count
   docker exec streamsre-postgres psql -U streamsre -c \
     "SELECT count(*) FROM pg_stat_activity;"
   ```

3. **Check for long-running queries**
   ```bash
   # TODO: Find slow queries
   docker exec streamsre-postgres psql -U streamsre -c \
     "SELECT pid, now() - pg_stat_activity.query_start AS duration, query
      FROM pg_stat_activity
      WHERE state = 'active' AND now() - pg_stat_activity.query_start > interval '5 seconds';"
   ```

4. **Check table sizes and bloat**
   ```bash
   # TODO: Check table sizes
   docker exec streamsre-postgres psql -U streamsre -c \
     "SELECT relname, pg_size_pretty(pg_total_relation_size(relid))
      FROM pg_catalog.pg_statio_user_tables ORDER BY pg_total_relation_size(relid) DESC;"
   ```

5. **Check for missing indexes**
   ```bash
   # TODO: Check index usage
   docker exec streamsre-postgres psql -U streamsre -c \
     "SELECT relname, seq_scan, idx_scan
      FROM pg_stat_user_tables WHERE seq_scan > 0;"
   ```

## Likely Causes

1. **Missing indexes** - Full table scans on large tables
2. **Connection pool exhaustion** - Too many concurrent connections
3. **Lock contention** - Multiple processes competing for same rows
4. **Table bloat** - Need VACUUM
5. **Resource exhaustion** - Postgres CPU/memory/disk I/O
6. **Network latency** - Connectivity issues between consumer and database

## Mitigation

### If queries are slow:
```bash
# TODO: Run EXPLAIN ANALYZE on slow queries
docker exec streamsre-postgres psql -U streamsre -c \
  "EXPLAIN ANALYZE SELECT * FROM reviews WHERE user_key = 'example';"

# TODO: Add missing indexes
docker exec streamsre-postgres psql -U streamsre -c \
  "CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_reviews_user_key ON reviews(user_key);"
```

### If table is bloated:
```bash
# TODO: Run VACUUM
docker exec streamsre-postgres psql -U streamsre -c "VACUUM ANALYZE reviews;"
```

### If connections are exhausted:
```bash
# TODO: Check max_connections setting
docker exec streamsre-postgres psql -U streamsre -c "SHOW max_connections;"

# TODO: Reduce consumer connection pool size or increase Postgres max_connections
```

### If there's lock contention:
```bash
# TODO: Identify blocking queries
docker exec streamsre-postgres psql -U streamsre -c \
  "SELECT blocked_locks.pid AS blocked_pid, blocked_activity.usename AS blocked_user,
          blocking_locks.pid AS blocking_pid, blocking_activity.usename AS blocking_user,
          blocked_activity.query AS blocked_statement
   FROM pg_catalog.pg_locks blocked_locks
   JOIN pg_catalog.pg_stat_activity blocked_activity ON blocked_activity.pid = blocked_locks.pid
   JOIN pg_catalog.pg_locks blocking_locks ON blocking_locks.locktype = blocked_locks.locktype
   JOIN pg_catalog.pg_stat_activity blocking_activity ON blocking_activity.pid = blocking_locks.pid
   WHERE NOT blocked_locks.granted;"
```

## Follow-ups

- [ ] Review query patterns and optimize if needed
- [ ] Set up connection pooling (e.g., PgBouncer) if not present
- [ ] Configure query timeout in consumer
- [ ] Add database performance alerts
- [ ] Consider read replicas for read-heavy workloads
- [ ] Schedule regular VACUUM jobs

