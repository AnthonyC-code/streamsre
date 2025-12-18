#!/bin/bash
# chaos_slow_consumer.sh - Simulate slow consumer processing
#
# This chaos script adds artificial latency to simulate a slow consumer.
# Useful for testing alerts and observing lag behavior.

set -euo pipefail

echo "=== Chaos Engineering: Slow Consumer ==="
echo ""

# TODO: Configuration
DELAY_MS=${1:-5000}  # Default 5 second delay
DURATION_SECONDS=${2:-60}  # Default 1 minute

echo "Simulating slow consumer with ${DELAY_MS}ms processing delay"
echo "Duration: ${DURATION_SECONDS} seconds"
echo ""

# TODO: Approach 1 - Use tc (traffic control) to add network latency
echo "TODO: Implement network latency injection"
echo "  Option 1: tc qdisc add - Add network delay"
echo "  Option 2: Set CHAOS_DELAY_MS environment variable"
echo "  Option 3: Use a chaos proxy"
echo ""

# TODO: Approach 2 - Environment variable that consumer respects
echo "Setting environment variable for consumer..."
echo "  export CHAOS_DELAY_MS=$DELAY_MS"
echo ""

# TODO: The consumer would need to check this variable:
# if os.Getenv("CHAOS_DELAY_MS") != "" {
#     delay, _ := strconv.Atoi(os.Getenv("CHAOS_DELAY_MS"))
#     time.Sleep(time.Duration(delay) * time.Millisecond)
# }

echo "Waiting ${DURATION_SECONDS} seconds..."
echo "TODO: Implement actual delay injection"
# sleep $DURATION_SECONDS

echo ""
echo "Cleanup (would remove the delay)..."
echo "TODO: Implement cleanup"
echo ""

echo "=== Chaos experiment complete ==="
echo ""
echo "Check Grafana for:"
echo "  - Increased consumer lag"
echo "  - Higher processing latency"
echo "  - Potential timeout errors"

