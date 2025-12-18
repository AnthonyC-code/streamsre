// Package db provides database connectivity using pgxpool.
//
// YOUR TASK (Milestone 3):
// Create a connection pool and implement basic operations.
package db

// TODO: Import required packages:
// - "context"
// - "github.com/jackc/pgx/v5/pgxpool"

// DB wraps a pgxpool.Pool for database operations.
//
// TODO: Define struct with field:
// - pool *pgxpool.Pool

// New creates a new database connection pool.
//
// TODO: Implement:
// 1. Call pgxpool.New(ctx, databaseURL)
// 2. Call pool.Ping(ctx) to verify connection
// 3. Return &DB{pool: pool}, nil on success
// func New(ctx context.Context, databaseURL string) (*DB, error)

// Close closes the database connection pool.
//
// TODO: Implement:
// - Call db.pool.Close()
// func (db *DB) Close()

// Ping verifies the database connection is alive.
//
// TODO: Implement:
// - Call db.pool.Ping(ctx)
// func (db *DB) Ping(ctx context.Context) error

// Pool returns the underlying pgxpool.Pool for direct access.
//
// TODO: Implement:
// - Return db.pool
// func (db *DB) Pool() *pgxpool.Pool
