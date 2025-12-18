// Package service provides HTTP endpoints.
//
// YOUR TASK (Milestone 11):
// Implement health check endpoints and metrics server.
package service

// TODO: Import:
// - "context"
// - "net/http"
// - "github.com/prometheus/client_golang/prometheus/promhttp"

// Server provides HTTP endpoints for health and metrics.
//
// TODO: Define struct with fields:
// - server  *http.Server
// - db      *db.DB       // For readiness check
// - ready   bool         // Are we ready for traffic?

// NewServer creates an HTTP server.
//
// TODO: Implement:
// 1. Create http.ServeMux
// 2. Register handlers:
//    - mux.HandleFunc("/healthz", s.handleHealthz)
//    - mux.HandleFunc("/readyz", s.handleReadyz)
//    - mux.Handle("/metrics", promhttp.Handler())
// 3. Create http.Server with mux
// func NewServer(addr string, db *db.DB) *Server

// Start starts the HTTP server in a goroutine.
//
// TODO: Implement:
// - go s.server.ListenAndServe()
// func (s *Server) Start() error

// Shutdown gracefully shuts down the server.
//
// TODO: Implement:
// - s.server.Shutdown(ctx)
// func (s *Server) Shutdown(ctx context.Context) error

// SetReady marks the server as ready/not ready.
// func (s *Server) SetReady(ready bool)

// handleHealthz handles liveness probes.
//
// LIVENESS: "Is the process alive?"
// - Always return 200 if we're running
// - Do NOT check dependencies (database, Kafka)
// - If this fails, Kubernetes RESTARTS the container
//
// TODO: Implement:
// - w.WriteHeader(http.StatusOK)
// - w.Write([]byte(`{"status":"ok"}`))
// func (s *Server) handleHealthz(w http.ResponseWriter, r *http.Request)

// handleReadyz handles readiness probes.
//
// READINESS: "Can this instance handle traffic?"
// - Check: Can we reach the database?
// - Check: Are we ready to consume? (s.ready)
// - If this fails, Kubernetes removes from load balancer (no traffic)
//   but does NOT restart
//
// TODO: Implement:
// 1. If !s.ready, return 503
// 2. If s.db.Ping(ctx) fails, return 503
// 3. Otherwise, return 200
// func (s *Server) handleReadyz(w http.ResponseWriter, r *http.Request)

// WHY SEPARATE LIVENESS AND READINESS?
//
// Scenario: Database is down
// - Liveness: 200 (process is fine!)
// - Readiness: 503 (can't do work right now)
// Result: Kubernetes stops sending traffic but doesn't restart.
//         When DB recovers, readiness returns 200, traffic resumes.
//
// Scenario: Process deadlock
// - Liveness: timeout/fail
// - Readiness: also fails
// Result: Kubernetes restarts the container.
//
// If liveness checked DB and DB was slow, Kubernetes would
// keep restarting healthy containers! That's why liveness
// should only check "is the process alive".
