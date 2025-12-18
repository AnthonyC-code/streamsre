// Package db provides database model definitions.
package db

import (
	"time"

	"github.com/google/uuid"
)

// Review represents a row in the reviews table.
type Review struct {
	// ReviewID is the primary key for the review.
	ReviewID string `db:"review_id"`

	// UserKey identifies the user who created the review.
	UserKey string `db:"user_key"`

	// Rating is the review rating (1-5).
	Rating int `db:"rating"`

	// Text is the review content.
	Text string `db:"text"`

	// CreatedAt is when the review was created.
	CreatedAt time.Time `db:"created_at"`
}

// ProcessedEvent represents a row in the processed_events table.
type ProcessedEvent struct {
	// EventID is the unique identifier for the processed event.
	EventID uuid.UUID `db:"event_id"`

	// ProcessedAt is when the event was processed.
	ProcessedAt time.Time `db:"processed_at"`
}

// DLQEventRecord represents a row in the dlq_events table.
type DLQEventRecord struct {
	// EventID is the unique identifier for the failed event.
	EventID uuid.UUID `db:"event_id"`

	// Reason describes why the event failed processing.
	Reason string `db:"reason"`

	// OriginalPayload is the raw JSON of the original event.
	OriginalPayload []byte `db:"original_payload"`

	// FailedAt is when the event was moved to the DLQ.
	FailedAt time.Time `db:"failed_at"`
}

