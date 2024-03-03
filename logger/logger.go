package logger

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/lvlcn-t/loggerhead/internal/logger"
)

// Logger is the interface for the logger.
// Its build on top of slog.Logger and extends it with additional logging methods.
type Logger = logger.Logger

// Options is the optional configuration for the logger.
type Options = logger.Options

// NewLogger creates a new Logger instance with optional configurations.
// The logger can be customized by passing an Options struct which allows for
// setting the log level, format, OpenTelemetry support, and a custom handler.
// If no Options are provided, default settings are applied based on environment variables or internal defaults.
//
// Example:
//
//	opts := logger.Options{Level: "INFO", Format: "JSON", OpenTelemetry: true}
//	log := logger.NewLogger(opts)
//	log.Info("Hello, world!")
func NewLogger(o ...logger.Options) logger.Logger {
	return logger.NewLogger(o...)
}

// NewNamedLogger creates a new Logger instance with the provided name and optional configurations.
// This function allows for the same level of customization as NewLogger, with the addition of setting a logger name.
//
// Example:
//
//	opts := logger.Options{Level: "DEBUG", Format: "TEXT"}
//	log := logger.NewNamedLogger("myServiceLogger", opts)
func NewNamedLogger(name string, o ...logger.Options) logger.Logger {
	return logger.NewNamedLogger(name, o...)
}

// NewContextWithLogger creates a new context based on the provided parent context.
// It embeds a logger into this new context, which is a child of the logger from the parent context.
// The child logger inherits settings from the parent.
// Returns the child context and its cancel function to cancel the new context.
//
// Note: If no logger is found in the parent context, a new logger with the default configuration is embedded into the new context.
func NewContextWithLogger(parent context.Context) (context.Context, context.CancelFunc) {
	return logger.NewContextWithLogger(parent)
}

// IntoContext embeds the provided slog.Logger into the given context and returns the modified context.
// This function is used for passing loggers through context, allowing for context-aware logging.
func IntoContext(ctx context.Context, log logger.Logger) context.Context {
	return logger.IntoContext(ctx, log)
}

// FromContext extracts the slog.Logger from the provided context.
// If the context does not have a logger, it returns a new logger with the default configuration.
// This function is useful for retrieving loggers from context in different parts of an application.
func FromContext(ctx context.Context) logger.Logger {
	return logger.FromContext(ctx)
}

// Middleware takes the logger from the context and adds it to the request context.
func Middleware(ctx context.Context) func(http.Handler) http.Handler {
	return logger.Middleware(ctx)
}

// FromSlog returns a new Logger instance from the provided slog.Logger.
func FromSlog(l *slog.Logger) logger.Logger {
	return logger.FromSlog(l)
}
