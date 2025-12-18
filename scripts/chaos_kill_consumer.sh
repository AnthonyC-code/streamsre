#!/bin/bash
# chaos_kill_consumer.sh - Simulate consumer crash
#
# This chaos script kills the consumer process to test:
# - Consumer restart behavior
# - Kafka consumer group rebalancing
# - Message reprocessing after crash

set -euo pipefail

echo "=== Chaos Engineering: Kill Consumer ==="
echo ""

# TODO: Configuration
RESTART_DELAY_SECONDS=${1:-30}  # Wait before restarting

echo "This script will:"
echo "  1. Find and kill the consumer process"
echo "  2. Wait ${RESTART_DELAY_SECONDS} seconds"
echo "  3. Observe the effects"
echo ""

# TODO: Find consumer process
echo "Finding consumer process..."
echo "TODO: Implement process discovery"
echo "  pgrep -f 'cmd/consumer' or similar"
echo ""

# TODO: Kill consumer
echo "Killing consumer..."
echo "TODO: Implement process kill"
echo "  kill -9 \$(pgrep -f 'cmd/consumer')"
echo ""

# TODO: Wait and observe
echo "Waiting ${RESTART_DELAY_SECONDS} seconds..."
echo "TODO: Implement wait period"
# sleep $RESTART_DELAY_SECONDS
echo ""

echo "=== Chaos experiment notes ==="
echo ""
echo "Expected behavior:"
echo "  - Consumer group will detect missing member"
echo "  - Partition rebalance will occur (if multiple consumers)"
echo "  - Uncommitted messages will be re-delivered"
echo "  - Lag will temporarily increase"
echo ""
echo "Monitor:"
echo "  - Consumer group status: rpk group describe streamsre-consumer"
echo "  - Consumer lag in Grafana"
echo "  - Duplicate message processing (idempotency check)"
echo ""
echo "Restart consumer with: make run-consumer"

