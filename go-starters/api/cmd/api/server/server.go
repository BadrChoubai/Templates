// Package server wires the HTTP server, its routes, middleware and database.
package server

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/badrchoubai/api/internal/routes"
	"github.com/badrchoubai/api/pkg/database"
	"github.com/badrchoubai/api/pkg/middleware"
	"github.com/badrchoubai/api/pkg/observability"
	"github.com/badrchoubai/api/pkg/observability/logging"
)

const (
	defaultHTTPPort         = "8084"
	defaultConnectionString = ".db/api.db"
)

var _ HTTPServer = (*Server)(nil)

// HTTPServer is the lifecycle the entry point depends on.
type HTTPServer interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

// Server is the application's HTTP server and the resources it owns.
type Server struct {
	httpServer *http.Server
	db         *sql.DB
	logger     *logging.Logger
}

// NewServer opens the database and configures the HTTP server from environment
// variables read through getEnv.
func NewServer(logger *logging.Logger, getEnv func(string) string) (*Server, error) {
	connStr := envOrDefault(getEnv, "DB_CONNECTION_STRING", defaultConnectionString)

	db, err := database.Open(connStr)
	if err != nil {
		return nil, err
	}

	port := envOrDefault(getEnv, "HTTP_PORT", defaultHTTPPort)
	srv := &Server{
		httpServer: &http.Server{
			Addr:              ":" + port,
			Handler:           setupHandler(logger, db),
			ReadHeaderTimeout: 5 * time.Second,
			ReadTimeout:       10 * time.Second,
			WriteTimeout:      10 * time.Second,
			IdleTimeout:       60 * time.Second,
		},
		db:     db,
		logger: logger,
	}

	return srv, nil
}

func envOrDefault(getEnv func(string) string, key, fallback string) string {
	if value := getEnv(key); value != "" {
		return value
	}
	return fallback
}

func setupHandler(logger *logging.Logger, db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, db)

	var handler http.Handler = mux
	handler = middleware.Heartbeat(handler, "/health")
	handler = observability.LogRequest(handler, logger)

	return handler
}

func addRoutes(mux *http.ServeMux, db *sql.DB) {
	routes.NewItemsRouter(mux, db)

	mux.Handle("/", routes.RootHandler())
}

// ListenAndServe starts accepting HTTP connections.
func (s *Server) ListenAndServe() error {
	s.logger.Info("starting HTTP server", "addr", "http://localhost"+s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully stops the HTTP server, then closes the database.
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("shutting down HTTP server")
	return errors.Join(s.httpServer.Shutdown(ctx), s.db.Close())
}
