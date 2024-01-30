package logger

import (
	"context"
	"log/slog"
	"net/http"
	"os"
)

type Logger interface {
	Core
	Debugf(msg string, args ...any)
	Infof(msg string, args ...any)
	Warnf(msg string, args ...any)
	Errorf(msg string, args ...any)
	Fatal(msg string, args ...any)
	Fatalf(msg string, args ...any)
	FatalContext(ctx context.Context, msg string, args ...any)
	Panic(msg string, args ...any)
	Panicf(msg string, args ...any)
	PanicContext(ctx context.Context, msg string, args ...any)
}

type logger struct {
	core Core
}

// NewLogger creates a new slog.Logger instance.
// If handlers are provided, the first handler in the slice is used; otherwise,
// a default JSON handler writing to os.Stderr is used. This function allows for
// custom configuration of logging handlers.
func NewLogger(h ...slog.Handler) Logger {
	return &logger{
		core: newCoreLogger(getHandler(h...)),
	}
}

// NewLogger creates a new slog.Logger instance.
// If handlers are provided, the first handler in the slice is used; otherwise,
// a default JSON handler writing to os.Stderr is used. This function allows for
// custom configuration of logging handlers.
// The loggers root group is the provided name.
func NewNamedLogger(name string, h ...slog.Handler) Logger {
	return &logger{
		core: With(newCoreLogger(getHandler(h...)), name),
	}
}

// NewContextWithLogger creates a new context based on the provided parent context.
// It embeds a logger into this new context, which is a child of the logger from the parent context.
// The child logger inherits settings from the parent.
// Returns the child context and its cancel function to cancel the new context.
func NewContextWithLogger(parent context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)
	return IntoContext(ctx, FromContext(parent)), cancel
}

// IntoContext embeds the provided slog.Logger into the given context and returns the modified context.
// This function is used for passing loggers through context, allowing for context-aware logging.
func IntoContext(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, logger{}, log)
}

// FromContext extracts the slog.Logger from the provided context.
// If the context does not have a logger, it returns a new logger with the default configuration.
// This function is useful for retrieving loggers from context in different parts of an application.
func FromContext(ctx context.Context) Logger {
	if ctx != nil {
		if logger, ok := ctx.Value(logger{}).(Logger); ok {
			return logger
		}
	}
	return NewLogger()
}

// Middleware takes the logger from the context and adds it to the request context
func Middleware(ctx context.Context) func(http.Handler) http.Handler {
	log := FromContext(ctx)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqCtx := IntoContext(r.Context(), log)
			next.ServeHTTP(w, r.WithContext(reqCtx))
		})
	}
}

func getHandler(h ...slog.Handler) slog.Handler {
	if len(h) > 0 {
		return h[0]
	}
	return slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     getLevel(os.Getenv("LOG_LEVEL")),
	})
}
