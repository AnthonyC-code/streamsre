#!/bin/bash
# chaos_slow_db.sh - Simulate slow database
#
# This chaos script adds artificial latency to Postgres queries to test:
# - Consumer retry behavior
# - Timeout handling
# - Circuit breaker patterns

set -euo pipefail

echo "=== Chaos Engineering: Slow Database ==="
echo ""

# TODO: Configuration
DELAY_MS=${1:-2000}  # Default 2 second delay per query
DURATION_SECONDS=${2:-60}  # Default 1 minute

echo "Simulating slow database with ${DELAY_MS}ms query delay"
echo "Duration: ${DURATION_SECONDS} seconds"
echo ""

POSTGRES_CONTAINER="streamsre-postgres"

# TODO: Option 1 - Use pg_sleep in a trigger
echo "Option 1: Add a trigger with pg_sleep"
echo "TODO: Implement trigger-based delay"
cat << 'EOF'
-- Create a trigger function that adds delay
CREATE OR REPLACE FUNCTION chaos_slow_insert()
RETURNS TRIGGER AS $$
BEGIN
    PERFORM pg_sleep(2);  -- 2 second delay
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Attach trigger to reviews table
CREATE TRIGGER chaos_reviews_slow
    BEFORE INSERT ON reviews
    FOR EACH ROW
    EXECUTE FUNCTION chaos_slow_insert();
EOF
echo ""

# TODO: Option 2 - Use tc to add network latency to Postgres port
echo "Option 2: Network latency with tc"
echo "TODO: Implement network-level delay"
echo "  tc qdisc add dev eth0 root netem delay ${DELAY_MS}ms"
echo ""

# TODO: Option 3 - Postgres statement timeout
echo "Option 3: Reduce statement_timeout to cause failures"
echo "TODO: Implement statement timeout reduction"
echo "  ALTER DATABASE streamsre SET statement_timeout = '100ms';"
echo ""

echo "Waiting ${DURATION_SECONDS} seconds..."
echo "TODO: Implement actual delay injection"
# sleep $DURATION_SECONDS
echo ""

echo "Cleanup..."
echo "TODO: Implement cleanup"
cat << 'EOF'
-- Remove chaos trigger
DROP TRIGGER IF EXISTS chaos_reviews_slow ON reviews;
DROP FUNCTION IF EXISTS chaos_slow_insert;

-- Reset statement timeout
ALTER DATABASE streamsre RESET statement_timeout;
EOF
echo ""

echo "=== Chaos experiment complete ==="
echo ""
echo "Check Grafana for:"
echo "  - Increased db_query_duration_seconds"
echo "  - Higher retry_count"
echo "  - Potential DLQ messages"
echo "  - Consumer lag increase"

