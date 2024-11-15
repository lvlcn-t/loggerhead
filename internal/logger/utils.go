package logger

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
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
func NewLogger(o ...Options) Provider {
	return &logger{
		Logger: slog.New(newHandler(o...)),
	}
}

// NewNamedLogger creates a new Logger instance with the provided name and optional configurations.
// This function allows for the same level of customization as NewLogger, with the addition of setting a logger name.
//
// Example:
//
//	opts := logger.Options{Level: "DEBUG", Format: "TEXT"}
//	log := logger.NewNamedLogger("myServiceLogger", opts)
func NewNamedLogger(name string, o ...Options) Provider {
	return &logger{
		Logger: slog.New(newHandler(o...)).With("name", name),
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

// ctxKey is the key used to store the logger in the context.
type ctxKey struct{}

// IntoContext embeds the provided slog.Logger into the given context and returns the modified context.
// This function is used for passing loggers through context, allowing for context-aware logging.
func IntoContext(ctx context.Context, log Provider) context.Context {
	return context.WithValue(ctx, ctxKey{}, log)
}

// FromContext extracts the slog.Logger from the provided context.
// If the context does not have a logger, it returns a new logger with the default configuration.
// This function is useful for retrieving loggers from context in different parts of an application.
func FromContext(ctx context.Context) Provider {
	if ctx != nil {
		if logger, ok := ctx.Value(ctxKey{}).(Provider); ok {
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

// ToSlog returns the underlying [slog.Logger].
func (l *logger) ToSlog() *slog.Logger {
	if l.Logger == nil {
		return slog.New(newHandler())
	}

	return l.Logger
}

// FromSlog returns a new Logger instance based on the provided [slog.Logger].
func FromSlog(l *slog.Logger) Provider {
	if l == nil {
		return NewLogger()
	}

	return &logger{l}
}

// newHandler returns a new slog.Handler based on the provided options.
//
// It returns the handler based on several conditions:
//  1. If a handler is provided, it returns the handler.
//  2. If OpenTelemetry support is enabled, it returns a new OtelHandler.
//  3. Otherwise, it returns a new BaseHandler.
func newHandler(o ...Options) slog.Handler {
	opts := newOptions(o...)
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
		log := clog.NewWithOptions(os.Stderr, clog.Options{
			TimeFormat:      time.Kitchen,
			Level:           clog.Level(newLevel(o.Level)),
			ReportTimestamp: true,
			ReportCaller:    true,
		})
		log.SetStyles(newCustomStyles())
		return log
	}

	return slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.Level(newLevel(o.Level)),
		ReplaceAttr: replaceAttr,
	})
}

// newCustomStyles returns the custom styles for the text logger.
func newCustomStyles() *clog.Styles {
	styles := clog.DefaultStyles()

	const maxWidth = 4
	for level, color := range LevelColors {
		styles.Levels[clog.Level(int(level))] = lipgloss.NewStyle().
			SetString(level.String()).
			Bold(true).
			MaxWidth(maxWidth).
			Foreground(lipgloss.Color(color))
	}

	return styles
}

// replaceAttr is the replacement function for slog.HandlerOptions.
func replaceAttr(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		lev := a.Value.Any().(slog.Level)
		a.Value = slog.StringValue(lev.String())
	}
	return a
}
