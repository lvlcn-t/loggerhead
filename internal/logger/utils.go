package logger

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	clog "github.com/charmbracelet/log"
)

// NewLogger creates a new Logger instance.
// If handlers are provided, the first handler in the slice is used; otherwise,
// a default JSON handler writing to os.Stderr is used. This function allows for
// custom configuration of logging handlers.
//
// Example:
//
//	log := logger.NewLogger()
//	log.Info("Hello, world!")
func NewLogger(h ...slog.Handler) Logger {
	return &logger{
		Logger: slog.New(getHandler(h...)),
	}
}

// NewNamedLogger creates a new Logger instance with the provided name.
// If handlers are provided, the first handler in the slice is used; otherwise,
// a default JSON handler writing to os.Stderr is used. This function allows for
// custom configuration of logging handlers.
func NewNamedLogger(name string, h ...slog.Handler) Logger {
	return &logger{
		Logger: slog.New(getHandler(h...)).With("name", name),
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

// ToSlog returns the underlying slog.Logger.
func (l *logger) ToSlog() *slog.Logger {
	return l.Logger
}

// FromSlog returns a new Logger instance based on the provided slog.Logger.
func FromSlog(l *slog.Logger) Logger {
	return &logger{l}
}

// getHandler returns the first handler in the slice if it exists; otherwise, it returns a new base handler.
func getHandler(h ...slog.Handler) slog.Handler {
	if len(h) > 0 {
		return h[0]
	}
	return newBaseHandler()
}

// newBaseHandler returns a new slog.Handler based on the environment variables.
func newBaseHandler() slog.Handler {
	l := getLevel(os.Getenv("LOG_LEVEL"))
	if strings.ToUpper(os.Getenv("LOG_FORMAT")) == "TEXT" {
		h := clog.New(os.Stderr)
		h.SetTimeFormat(time.Kitchen)
		h.SetReportTimestamp(true)
		h.SetReportCaller(true)
		h.SetLevel(clog.Level(l))
		return h
	}

	return slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource:   true,
		Level:       Level(l),
		ReplaceAttr: replaceAttr,
	})
}

// replaceAttr is the replacement function for slog.HandlerOptions.
func replaceAttr(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		lev := a.Value.Any().(Level)
		a.Value = slog.StringValue(getLevelString(lev))
	}
	return a
}
