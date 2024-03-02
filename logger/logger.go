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

// NewLogger creates a new Logger instance.
// If handlers are provided, the first handler in the slice is used; otherwise,
// a default JSON handler writing to os.Stderr is used. This function allows for
// custom configuration of logging handlers.
//
// Example:
//
//	log := logger.NewLogger()
//	log.Info("Hello, world!")
func NewLogger(h ...slog.Handler) logger.Logger {
	return logger.NewLogger(h...)
}

// NewNamedLogger creates a new Logger instance with the provided name.
// If handlers are provided, the first handler in the slice is used; otherwise,
// a default JSON handler writing to os.Stderr is used. This function allows for
// custom configuration of logging handlers.
func NewNamedLogger(name string, h ...slog.Handler) logger.Logger {
	return logger.NewNamedLogger(name, h...)
}

// NewContextWithLogger creates a new context based on the provided parent context.
// It embeds a logger into this new context, which is a child of the logger from the parent context.
// The child logger inherits settings from the parent.
// Returns the child context and its cancel function to cancel the new context.
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
