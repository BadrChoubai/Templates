// Package database opens and configures the application's SQLite connection pool.
package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3" // registers the sqlite3 driver
)

const (
	maxOpenConns = 5
	maxIdleConns = 2
	maxLifetime  = 2 * time.Minute
	maxIdleTime  = 20 * time.Second
)

// Open opens a SQLite database at connectionStr and verifies the connection.
func Open(connectionStr string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	db.SetConnMaxLifetime(maxLifetime)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("connecting to database: %w", err)
	}

	return db, nil
}
