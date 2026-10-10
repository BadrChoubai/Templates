package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/badrchoubai/api/cmd/api/server"
	"github.com/badrchoubai/api/pkg/observability/logging"
)

func main() {
	ctx := context.Background()

	if err := run(ctx, os.Stdout, os.Getenv); err != nil {
		os.Exit(1)
	}
}

func run(ctx context.Context, stdout io.Writer, getenv func(string) string) error {
	var wg sync.WaitGroup

	logger := logging.NewLogger(stdout)
	srv, err := server.NewServer(logger, getenv)
	if err != nil {
		return err
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		logger.Info(fmt.Sprintf("caught signal: %s", s))

		shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			shutdownError <- err
		}

		logger.Info("completing background tasks")

		wg.Wait()
		shutdownError <- nil
	}()

	err = srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	return nil
}
