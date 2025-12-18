#!/bin/bash
# seed_topics.sh - Create Kafka topics for the StreamSRE project
#
# Run this after starting infrastructure with `make up`.
# This is the ONLY script you should run as-is (it creates necessary topics).

set -euo pipefail

echo "=== StreamSRE Topic Seeding ==="
echo ""

REDPANDA_CONTAINER="streamsre-redpanda"
MAIN_TOPIC="events.main"
DLQ_TOPIC="events.dlq"
PARTITIONS=6
REPLICATION_FACTOR=1

echo "Waiting for Redpanda to be ready..."
until docker exec $REDPANDA_CONTAINER rpk cluster health 2>/dev/null | grep -q "Healthy"; do
    echo "  Redpanda not ready yet, waiting..."
    sleep 2
done
echo "Redpanda is ready!"
echo ""

echo "Creating topic: $MAIN_TOPIC with $PARTITIONS partitions..."
docker exec $REDPANDA_CONTAINER rpk topic create $MAIN_TOPIC \
    --partitions $PARTITIONS \
    --replicas $REPLICATION_FACTOR \
    2>/dev/null || echo "  (Topic may already exist)"

echo "Creating topic: $DLQ_TOPIC..."
docker exec $REDPANDA_CONTAINER rpk topic create $DLQ_TOPIC \
    --partitions 1 \
    --replicas $REPLICATION_FACTOR \
    2>/dev/null || echo "  (Topic may already exist)"

echo ""
echo "Listing topics:"
docker exec $REDPANDA_CONTAINER rpk topic list

echo ""
echo "=== Topic seeding complete ==="
echo ""
echo "You can now run the producer and consumer."
