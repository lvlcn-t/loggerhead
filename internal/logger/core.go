package logger

import (
	"context"
	"log/slog"
)

var _ Core = (*coreLogger)(nil)

type Core interface {
	// Debug logs at LevelDebug.
	Debug(msg string, args ...any)
	// DebugContext logs at LevelDebug with the given context.
	DebugContext(ctx context.Context, msg string, args ...any)
	// Info logs at LevelInfo.
	Info(msg string, args ...any)
	// InfoContext logs at LevelInfo with the given context.
	InfoContext(ctx context.Context, msg string, args ...any)
	// Warn logs at LevelWarn.
	Warn(msg string, args ...any)
	// WarnContext logs at LevelWarn with the given context.
	WarnContext(ctx context.Context, msg string, args ...any)
	// Error logs at LevelError.
	Error(msg string, args ...any)
	// ErrorContext logs at LevelError with the given context.
	ErrorContext(ctx context.Context, msg string, args ...any)
	// With calls Logger.With on the default logger.
	With(args ...any) *slog.Logger
	// WithGroup returns a Logger that starts a group, if name is non-empty.
	// The keys of all attributes added to the Logger will be qualified by the given
	// name. (How that qualification happens depends on the [Handler.WithGroup]
	// method of the Logger's Handler.)
	//
	// If name is empty, WithGroup returns the receiver.
	WithGroup(name string) *slog.Logger
	// Log emits a log record with the current time and the given level and message.
	// The Record's Attrs consist of the Logger's attributes followed by
	// the Attrs specified by args.
	//
	// The attribute arguments are processed as follows:
	//   - If an argument is an Attr, it is used as is.
	//   - If an argument is a string and this is not the last argument,
	//     the following argument is treated as the value and the two are combined
	//     into an Attr.
	//   - Otherwise, the argument is treated as a value with key "!BADKEY".
	Log(ctx context.Context, level Level, msg string, args ...any)
	// LogAttrs is a more efficient version of [Logger.Log] that accepts only Attrs.
	LogAttrs(ctx context.Context, level Level, msg string, attrs ...slog.Attr)
	// Handler returns l's Handler.
	Handler() slog.Handler
	// Enabled reports whether l emits log records at the given context and level.
	Enabled(ctx context.Context, level Level) bool
}

type coreLogger struct {
	*slog.Logger
}

func newCoreLogger(h slog.Handler) *coreLogger {
	return &coreLogger{
		slog.New(h),
	}
}

func With(l Core, args ...any) *coreLogger {
	return &coreLogger{
		l.With(args...),
	}
}

func WithGroup(l Core, name string) *coreLogger {
	return &coreLogger{
		l.WithGroup(name),
	}
}
