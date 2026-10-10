package server

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/badrchoubai/api/internal/routes"
	"github.com/badrchoubai/api/pkg/database"
	"github.com/badrchoubai/api/pkg/middleware"
	"github.com/badrchoubai/api/pkg/observability"
	"github.com/badrchoubai/api/pkg/observability/logging"
)

var _ HTTPServer = (*Server)(nil)

type HTTPServer interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

type Server struct {
	HttpServer *http.Server
	logger     *logging.Logger
}

func NewServer(logger *logging.Logger, getEnv func(string) string) (*Server, error) {
	connStr := getEnv("DB_CONNECTION_STRING")

	db, err := database.Open(connStr)
	if err != nil {
		logger.Warn(err.Error())
		return nil, err
	}

	handler := setupHandler(logger, db)

	port := getEnv("HTTP_PORT")
	srv := &Server{
		HttpServer: &http.Server{
			Addr:    ":" + port,
			Handler: handler,
		},
		logger: logger,
	}

	return srv, nil
}

func setupHandler(logger *logging.Logger, db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, db)

	var handler http.Handler = mux
	handler = middleware.Heartbeat(handler, "/health")
	handler = observability.LogRequest(handler, logger)

	return handler
}

func addRoutes(mux *http.ServeMux, database *sql.DB) {
	routes.NewItemsRouter(mux, database)

	mux.Handle("/", routes.RootHandler())
}

func (s *Server) ListenAndServe() error {
	s.logger.Info("Starting HTTP Server")
	s.logger.Info("http://localhost" + s.HttpServer.Addr)
	return s.HttpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down HTTP Server")
	return s.HttpServer.Shutdown(ctx)
}
