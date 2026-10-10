package logging

import (
	"io"
	"log/slog"
)

type StructuredLogger interface {
	Info(msg string, args ...any)
	Error(whatWasHappening string, args ...any)
	Fatal(msg string, args ...any)
	Warn(msg string, args ...any)
	Debug(msg string, args ...any)
}

type Logger struct {
	log *slog.Logger
}

func NewLogger(out io.Writer) *Logger {
	handler := createHandler(out)

	return &Logger{
		log: slog.New(handler),
	}
}

func createHandler(out io.Writer) *slog.JSONHandler {
	return slog.NewJSONHandler(out, nil)
}

func (l *Logger) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
}

func (l *Logger) Error(whatWasHappening string, args ...any) {
	l.log.Error(whatWasHappening, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.log.Warn(msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	l.log.Debug(msg, args...)
}
