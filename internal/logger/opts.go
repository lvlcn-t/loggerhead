package logger

import (
	"log/slog"
	"os"
)

// Options is the optional configuration for the logger.
type Options struct {
	// Level is the minimum log level.
	Level string
	// Format is the log format.
	Format string
	// OpenTelemetry is a flag to enable OpenTelemetry support.
	OpenTelemetry bool
	// Handler is the log handler.
	Handler slog.Handler
}

// newDefaultOptions returns the default Options.
func newDefaultOptions() Options {
	return Options{
		Level:         os.Getenv("LOG_LEVEL"),
		Format:        os.Getenv("LOG_FORMAT"),
		OpenTelemetry: false,
	}
}

// newOptions creates a new Options instance with the provided Options merged with the default Options.
func newOptions(o ...Options) Options {
	opts := newDefaultOptions()
	if len(o) > 0 {
		return o[0].merge(opts)
	}
	return opts
}

// merge merges the provided Options with the receiver Options.
func (o *Options) merge(d Options) Options {
	_, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		d.Level = o.Level
	}
	_, ok = os.LookupEnv("LOG_FORMAT")
	if !ok {
		d.Format = o.Format
	}
	if o.OpenTelemetry {
		d.OpenTelemetry = o.OpenTelemetry
	}
	if o.Handler != nil {
		d.Handler = o.Handler
	}
	return d
}
