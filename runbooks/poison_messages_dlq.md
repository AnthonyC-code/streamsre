# Runbook: Poison Messages in DLQ

## Symptoms

- Prometheus alert: `DLQMessagesHigh`
- `events_dlq_total` metric increasing
- `dlq_events` table has new entries
- Consumer logs show repeated processing failures for specific events

## Immediate Checks

1. **Check DLQ event count**
   ```bash
   # TODO: Query DLQ table count
   docker exec streamsre-postgres psql -U streamsre -c \
     "SELECT COUNT(*) FROM dlq_events;"
   ```

2. **Sample recent DLQ events**
   ```bash
   # TODO: View recent DLQ entries
   docker exec streamsre-postgres psql -U streamsre -c \
     "SELECT event_id, reason, failed_at FROM dlq_events ORDER BY failed_at DESC LIMIT 10;"
   ```

3. **Check DLQ event payloads**
   ```bash
   # TODO: Inspect specific DLQ event payload
   docker exec streamsre-postgres psql -U streamsre -c \
     "SELECT event_id, reason, original_payload FROM dlq_events ORDER BY failed_at DESC LIMIT 5;"
   ```

4. **Check consumer logs for error patterns**
   ```bash
   # TODO: Search for error patterns
   make logs | grep -E "(ERROR|WARN|failed|poison)"
   ```

5. **Check DLQ Kafka topic (if using Kafka DLQ)**
   ```bash
   # TODO: Check DLQ topic
   docker exec streamsre-redpanda rpk topic consume reviews.dlq --num 5
   ```

## Likely Causes

1. **Malformed event data** - Invalid JSON, missing required fields
2. **Schema mismatch** - Producer sending events with unexpected schema
3. **Business rule violations** - Invalid rating values, duplicate IDs
4. **Encoding issues** - Character encoding problems
5. **Producer bug** - Faulty event generation logic
6. **Database constraint violations** - Data doesn't satisfy table constraints

## Mitigation

### Analyze the poison messages:
```bash
# TODO: Group DLQ events by reason
docker exec streamsre-postgres psql -U streamsre -c \
  "SELECT reason, COUNT(*) as count FROM dlq_events GROUP BY reason ORDER BY count DESC;"
```

### If schema mismatch:
```bash
# TODO: Compare event schema with expected format
# TODO: Check producer version and update if needed
# TODO: Add schema validation at producer
```

### If data is recoverable:
```bash
# TODO: Fix data and replay
# - Extract event from DLQ
# - Fix the payload
# - Re-publish to main topic or process directly
```

### If producer is sending bad data:
```bash
# TODO: Identify and fix producer issue
# TODO: Consider adding event validation at producer side
# TODO: Add schema registry for event validation
```

### Clear processed DLQ events:
```bash
# TODO: After investigation, clear old DLQ events
docker exec streamsre-postgres psql -U streamsre -c \
  "DELETE FROM dlq_events WHERE failed_at < NOW() - INTERVAL '7 days';"
```

## Follow-ups

- [ ] Identify root cause of poison messages
- [ ] Fix producer or consumer code as needed
- [ ] Add input validation before processing
- [ ] Consider implementing schema validation
- [ ] Set up DLQ replay mechanism for recoverable events
- [ ] Create alert for DLQ growth rate
- [ ] Document common DLQ reasons and fixes
- [ ] Review event schema and add versioning if needed

