#!/bin/bash
# seed_topics.sh - Create Kafka topics for the StreamSRE project
#
# This script creates the required topics in Redpanda/Kafka.
# Run after starting the infrastructure with `make up`.

set -euo pipefail

echo "=== StreamSRE Topic Seeding ==="
echo ""

# TODO: Configuration - update these as needed
REDPANDA_CONTAINER="streamsre-redpanda"
MAIN_TOPIC="reviews"
DLQ_TOPIC="reviews.dlq"
PARTITIONS=3
REPLICATION_FACTOR=1

echo "Creating topics in Redpanda..."
echo ""

# TODO: Create main topic
echo "Creating topic: $MAIN_TOPIC"
# docker exec $REDPANDA_CONTAINER rpk topic create $MAIN_TOPIC \
#   --partitions $PARTITIONS \
#   --replicas $REPLICATION_FACTOR

echo "TODO: Implement topic creation"
echo "  rpk topic create $MAIN_TOPIC --partitions $PARTITIONS --replicas $REPLICATION_FACTOR"
echo ""

# TODO: Create DLQ topic
echo "Creating topic: $DLQ_TOPIC"
# docker exec $REDPANDA_CONTAINER rpk topic create $DLQ_TOPIC \
#   --partitions $PARTITIONS \
#   --replicas $REPLICATION_FACTOR

echo "TODO: Implement DLQ topic creation"
echo "  rpk topic create $DLQ_TOPIC --partitions $PARTITIONS --replicas $REPLICATION_FACTOR"
echo ""

# TODO: List topics to verify
echo "Listing topics..."
# docker exec $REDPANDA_CONTAINER rpk topic list

echo "TODO: Implement topic listing"
echo "  rpk topic list"
echo ""

echo "=== Topic seeding complete ==="

