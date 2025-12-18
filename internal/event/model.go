// Package event defines the event models used throughout the system.
//
// YOUR TASK (Milestone 1):
// Define the Event struct that represents messages in Kafka.
// See SPEC.md for the exact schema.
package event

// TODO: Import required packages
// - "time"
// - "github.com/google/uuid"

// Event represents a single message in Kafka.
//
// TODO: Define struct with these fields:
// - EventID       uuid.UUID   `json:"event_id"`       // For idempotency
// - EventType     string      `json:"event_type"`     // e.g., "review_created"
// - Key           string      `json:"key"`            // Kafka partition key
// - Timestamp     time.Time   `json:"ts"`             // When created
// - SchemaVersion int         `json:"schema_version"` // For versioning
// - Payload       ReviewData  `json:"payload"`        // The actual data

// ReviewData is the payload for review events.
//
// TODO: Define struct with these fields:
// - ReviewID string `json:"review_id"`
// - Rating   int    `json:"rating"`
// - Text     string `json:"text"`

// Validate checks if the event is valid.
//
// TODO: Implement validation:
// - EventID must not be zero
// - EventType must not be empty
// - Key must not be empty
// - Rating must be 1-5
// - Return error describing what's invalid
