// Package event defines the event models used throughout the system.
// Events are the core data structures exchanged between producer and consumer.
package event

import (
	"time"

	"github.com/google/uuid"
)

// ReviewEvent represents a product review event published to Kafka.
type ReviewEvent struct {
	// EventID is a unique identifier for this event (used for idempotency).
	EventID uuid.UUID `json:"event_id"`

	// ReviewID is the business identifier for the review.
	ReviewID string `json:"review_id"`

	// UserKey identifies the user who created the review.
	UserKey string `json:"user_key"`

	// Rating is the review rating (1-5).
	Rating int `json:"rating"`

	// Text is the review content.
	Text string `json:"text"`

	// CreatedAt is when the review was created.
	CreatedAt time.Time `json:"created_at"`
}

// Validate checks if the event has all required fields and valid values.
func (e *ReviewEvent) Validate() error {
	// TODO: Implement validation logic
	// - Check EventID is not zero
	// - Check ReviewID is not empty
	// - Check UserKey is not empty
	// - Check Rating is between 1 and 5
	// - Check Text is not empty
	// - Check CreatedAt is not zero
	panic("TODO")
}

// DLQEvent represents an event that failed processing and was sent to the DLQ.
type DLQEvent struct {
	// EventID is the original event's unique identifier.
	EventID uuid.UUID `json:"event_id"`

	// Reason describes why the event failed processing.
	Reason string `json:"reason"`

	// OriginalPayload is the raw JSON of the original event.
	OriginalPayload []byte `json:"original_payload"`

	// FailedAt is when the event was moved to the DLQ.
	FailedAt time.Time `json:"failed_at"`
}

