// Package db provides SQL query functions.
//
// YOUR TASK (Milestone 3):
// Implement the database queries needed for event processing.
package db

// TODO: Import required packages:
// - "context"
// - "github.com/google/uuid"

// Queries provides database query operations.
//
// TODO: Define struct with field:
// - db *DB

// NewQueries creates a new Queries instance.
// func NewQueries(db *DB) *Queries

// IsEventProcessed checks if an event has already been processed.
//
// TODO: Implement:
// 1. Query: SELECT 1 FROM processed_events WHERE event_id = $1
// 2. If row exists, return true, nil
// 3. If no row, return false, nil
// 4. Handle pgx.ErrNoRows appropriately
// func (q *Queries) IsEventProcessed(ctx context.Context, eventID uuid.UUID) (bool, error)

// MarkEventProcessed records that an event has been processed.
//
// TODO: Implement:
// - Query: INSERT INTO processed_events (event_id) VALUES ($1)
// func (q *Queries) MarkEventProcessed(ctx context.Context, eventID uuid.UUID) error

// InsertReview inserts a new review into the database.
//
// TODO: Implement:
// - Query: INSERT INTO reviews (review_id, user_key, rating, text, created_at) VALUES ($1, $2, $3, $4, $5)
// func (q *Queries) InsertReview(ctx context.Context, r *Review) error

// InsertDLQEvent records a failed event in the DLQ table.
//
// TODO: Implement:
// - Query: INSERT INTO dlq_events (event_id, reason, original_payload) VALUES ($1, $2, $3)
// func (q *Queries) InsertDLQEvent(ctx context.Context, eventID uuid.UUID, reason string, payload []byte) error

// ProcessEventTx handles event processing in a SINGLE TRANSACTION.
//
// THIS IS THE MOST IMPORTANT FUNCTION - IT ENSURES EXACTLY-ONCE SEMANTICS!
//
// TODO: Implement:
// 1. Begin transaction: tx, err := db.pool.Begin(ctx)
// 2. defer tx.Rollback(ctx)  // Always rollback if not committed
// 3. Check if event already processed:
//    - SELECT 1 FROM processed_events WHERE event_id = $1
//    - If exists: return nil (already done, skip!)
// 4. Insert the review:
//    - INSERT INTO reviews (...) VALUES (...)
// 5. Mark event as processed:
//    - INSERT INTO processed_events (event_id) VALUES ($1)
// 6. Commit: return tx.Commit(ctx)
//
// WHY TRANSACTIONAL?
// If you insert review but crash before marking processed,
// on restart you'd insert the review AGAIN (duplicate!).
// The transaction ensures both happen or neither happens.
// func (q *Queries) ProcessEventTx(ctx context.Context, eventID uuid.UUID, r *Review) error
