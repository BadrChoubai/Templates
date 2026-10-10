// Package logging provides a structured JSON logger built on log/slog.
package logging

import (
	"io"
	"log/slog"
)

var _ StructuredLogger = (*Logger)(nil)

// StructuredLogger logs a message with alternating key/value pairs.
type StructuredLogger interface {
	Info(msg string, args ...any)
	Error(whatWasHappening string, args ...any)
	Warn(msg string, args ...any)
	Debug(msg string, args ...any)
}

// Logger writes structured JSON logs.
type Logger struct {
	log *slog.Logger
}

// NewLogger returns a Logger that writes JSON lines to out.
func NewLogger(out io.Writer) *Logger {
	return &Logger{
		log: slog.New(slog.NewJSONHandler(out, nil)),
	}
}

// Info logs at info level.
func (l *Logger) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
}

// Error logs at error level.
func (l *Logger) Error(whatWasHappening string, args ...any) {
	l.log.Error(whatWasHappening, args...)
}

// Warn logs at warn level.
func (l *Logger) Warn(msg string, args ...any) {
	l.log.Warn(msg, args...)
}

// Debug logs at debug level.
func (l *Logger) Debug(msg string, args ...any) {
	l.log.Debug(msg, args...)
}
