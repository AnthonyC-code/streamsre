// Package db provides database model definitions.
//
// YOUR TASK (Milestone 3):
// Define structs that map to your database tables.
package db

// TODO: Import required packages:
// - "time"
// - "github.com/google/uuid"

// Review represents a row in the reviews table.
//
// TODO: Define struct with fields:
// - ReviewID  string    // Primary key
// - UserKey   string    // Who wrote it (from event Key)
// - Rating    int       // 1-5
// - Text      string    // Review content
// - CreatedAt time.Time // When it was created

// ProcessedEvent represents a row in the processed_events table.
//
// TODO: Define struct with fields:
// - EventID     uuid.UUID
// - ProcessedAt time.Time

// DLQEvent represents a row in the dlq_events table.
//
// TODO: Define struct with fields:
// - EventID         uuid.UUID
// - Reason          string    // Why it failed
// - OriginalPayload []byte    // Raw JSON for debugging
// - FailedAt        time.Time
