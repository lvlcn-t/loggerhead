package logger

import (
	"log/slog"
	"os"
)

// Opts is the optional configuration for the logger.
type Opts struct {
	// Level is the minimum log level.
	Level string
	// Format is the log format.
	Format string
	// OpenTelemetry is a flag to enable OpenTelemetry support.
	OpenTelemetry bool
	// Handler is the log handler.
	Handler slog.Handler
}

// newDefaultOpts returns the default Opts.
func newDefaultOpts() Opts {
	return Opts{
		Level:         os.Getenv("LOG_LEVEL"),
		Format:        os.Getenv("LOG_FORMAT"),
		OpenTelemetry: false,
	}
}

// getOpts returns the first Opts in the slice if it exists; otherwise, it returns the default Opts.
func getOpts(o ...Opts) Opts {
	opts := newDefaultOpts()
	if len(o) > 0 {
		return o[0].merge(opts)
	}
	return opts
}

// merge merges the provided Opts with the receiver Opts.
func (o Opts) merge(d Opts) Opts {
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
