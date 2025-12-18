// Package service provides HTTP handlers for metrics and health checks.
package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	"streamsre/internal/db"
)

// HTTPServer provides HTTP endpoints for metrics and health checks.
type HTTPServer struct {
	server *http.Server
	db     *db.DB
	logger *zap.Logger
	ready  bool
}

// NewHTTPServer creates a new HTTP server for metrics and health.
func NewHTTPServer(port int, database *db.DB, logger *zap.Logger) *HTTPServer {
	s := &HTTPServer{
		db:     database,
		logger: logger,
		ready:  false,
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/healthz", s.handleHealthz)
	mux.HandleFunc("/readyz", s.handleReadyz)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	return s
}

// Start starts the HTTP server.
func (s *HTTPServer) Start() error {
	// TODO: Start server in background
	// TODO: Log startup message
	panic("TODO")
}

// Shutdown gracefully shuts down the HTTP server.
func (s *HTTPServer) Shutdown(ctx context.Context) error {
	// TODO: Shutdown server with context
	panic("TODO")
}

// SetReady marks the server as ready to receive traffic.
func (s *HTTPServer) SetReady(ready bool) {
	s.ready = ready
}

// handleHealthz handles liveness probe requests.
func (s *HTTPServer) handleHealthz(w http.ResponseWriter, r *http.Request) {
	// TODO: Return 200 OK if server is alive
	// Liveness just checks if the process is running
	panic("TODO")
}

// handleReadyz handles readiness probe requests.
func (s *HTTPServer) handleReadyz(w http.ResponseWriter, r *http.Request) {
	// TODO: Check if server is ready to receive traffic
	// - Check database connection
	// - Check if consumer is ready
	// Return 200 if ready, 503 if not
	panic("TODO")
}

