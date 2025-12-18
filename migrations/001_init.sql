-- StreamSRE Database Schema
-- This migration creates the core tables for event processing

-- Idempotency table: tracks which events have been processed
-- Used to prevent duplicate processing of the same event
CREATE TABLE IF NOT EXISTS processed_events (
    event_id UUID PRIMARY KEY,
    processed_at TIMESTAMPTZ DEFAULT NOW()
);

-- Reviews table: stores the actual review data from events
CREATE TABLE IF NOT EXISTS reviews (
    review_id TEXT PRIMARY KEY,
    user_key TEXT NOT NULL,
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    text TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);

-- Dead Letter Queue table: stores events that failed processing
-- Allows for later investigation and replay
CREATE TABLE IF NOT EXISTS dlq_events (
    event_id UUID PRIMARY KEY,
    reason TEXT NOT NULL,
    original_payload JSONB NOT NULL,
    failed_at TIMESTAMPTZ DEFAULT NOW()
);

-- Indexes for common query patterns
CREATE INDEX IF NOT EXISTS idx_processed_events_processed_at ON processed_events(processed_at);
CREATE INDEX IF NOT EXISTS idx_reviews_user_key ON reviews(user_key);
CREATE INDEX IF NOT EXISTS idx_reviews_created_at ON reviews(created_at);
CREATE INDEX IF NOT EXISTS idx_dlq_events_failed_at ON dlq_events(failed_at);

