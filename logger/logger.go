package logger

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/lvlcn-t/loggerhead/internal/logger"
)

// Logger is an alias for the [Provider] interface.
// It is defined for backward compatibility with previous versions of the logger package.
//
// Deprecated: Use [Provider] instead.
type Logger = Provider // TODO: Remove in v0.4.0

// Provider is the interface for the logger.
// Its build on top of slog.Logger and extends it with additional logging methods.
type Provider = logger.Provider

// Options is the optional configuration for the logger.
type Options = logger.Options

// Level is a custom type for log levels.
type Level = logger.Level

const (
	// LevelTrace represents the TRACE log level.
	LevelTrace = logger.LevelTrace
	// LevelDebug represents the DEBUG log level.
	LevelDebug = logger.LevelDebug
	// LevelInfo represents the INFO log level.
	LevelInfo = logger.LevelInfo
	// LevelNotice represents the NOTICE log level.
	LevelNotice = logger.LevelNotice
	// LevelWarn represents the WARN log level.
	LevelWarn = logger.LevelWarn
	// LevelError represents the ERROR log level.
	LevelError = logger.LevelError
	// LevelPanic represents the PANIC log level.
	LevelPanic = logger.LevelPanic
	// LevelFatal represents the FATAL log level.
	LevelFatal = logger.LevelFatal
)

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
func NewLogger(o ...logger.Options) logger.Provider {
	return logger.NewLogger(o...)
}

// NewNamedLogger creates a new Logger instance with the provided name and optional configurations.
// This function allows for the same level of customization as NewLogger, with the addition of setting a logger name.
//
// Example:
//
//	opts := logger.Options{Level: "DEBUG", Format: "TEXT"}
//	log := logger.NewNamedLogger("myServiceLogger", opts)
func NewNamedLogger(name string, o ...logger.Options) logger.Provider {
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

// IntoContext embeds the provided [logger.Provider] into the given context and returns the modified context.
// This function is used for passing loggers through context, allowing for context-aware logging.
func IntoContext(ctx context.Context, log logger.Provider) context.Context {
	return logger.IntoContext(ctx, log)
}

// FromContext extracts the [logger.Provider] from the provided context.
// If the context does not have a logger, it returns a new logger with the default configuration.
// This function is useful for retrieving loggers from context in different parts of an application.
func FromContext(ctx context.Context) logger.Provider {
	return logger.FromContext(ctx)
}

// Middleware takes the logger from the context and adds it to the request context.
func Middleware(ctx context.Context) func(http.Handler) http.Handler {
	return logger.Middleware(ctx)
}

// FromSlog returns a new [Logger] instance from the provided [slog.Logger].
func FromSlog(l *slog.Logger) logger.Provider {
	return logger.FromSlog(l)
}
