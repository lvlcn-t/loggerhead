package logger

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	clog "github.com/charmbracelet/log"
	otel "github.com/remychantenay/slog-otel"
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
func NewLogger(o ...Options) Logger {
	return &logger{
		Logger: slog.New(getHandler(o...)),
	}
}

// NewNamedLogger creates a new Logger instance with the provided name and optional configurations.
// This function allows for the same level of customization as NewLogger, with the addition of setting a logger name.
//
// Example:
//
//	opts := logger.Options{Level: "DEBUG", Format: "TEXT"}
//	log := logger.NewNamedLogger("myServiceLogger", opts)
func NewNamedLogger(name string, o ...Options) Logger {
	return &logger{
		Logger: slog.New(getHandler(o...)).With("name", name),
	}
}

// NewContextWithLogger creates a new context based on the provided parent context.
// It embeds a logger into this new context, which is a child of the logger from the parent context.
// The child logger inherits settings from the parent.
// Returns the child context and its cancel function to cancel the new context.
func NewContextWithLogger(ctx context.Context) (context.Context, context.CancelFunc) {
	c, cancel := context.WithCancel(ctx)
	return IntoContext(c, FromContext(ctx)), cancel
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

// getHandler returns a new slog.Handler based on the provided options.
//
// It returns the handler based on several conditions:
//  1. If a handler is provided, it returns the handler.
//  2. If OpenTelemetry support is enabled, it returns a new OtelHandler.
//  3. Otherwise, it returns a new BaseHandler.
func getHandler(o ...Options) slog.Handler {
	opts := getOptions(o...)
	if opts.Handler != nil {
		return opts.Handler
	}
	if opts.OpenTelemetry {
		return otel.NewOtelHandler()(newBaseHandler(opts))
	}
	return newBaseHandler(opts)
}

// newBaseHandler returns a new slog.Handler based on the environment variables.
func newBaseHandler(o Options) slog.Handler {
	if strings.EqualFold(o.Format, "TEXT") {
		// TODO: customize level output as soon as clog v0.4.0 is released: https://github.com/charmbracelet/log/issues/88#issuecomment-2000161131
		return clog.NewWithOptions(os.Stderr, clog.Options{
			TimeFormat:      time.Kitchen,
			Level:           clog.Level(getLevel(o.Level)),
			ReportTimestamp: true,
			ReportCaller:    true,
		})
	}

	return slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource:   true,
		Level:       Level(getLevel(o.Level)),
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
