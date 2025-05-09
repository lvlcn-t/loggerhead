package logger

import (
	"context"
	"log/slog"
	"runtime"
	"time"
)

var _ Provider = (*logger)(nil)

// Provider is a interface that provides logging methods.
type Provider interface {
	// Trace logs at [LevelTrace].
	Trace(msg string, args ...any)
	// Tracef logs at [LevelTrace].
	// Arguments are handled in the manner of [fmt.Printf].
	Tracef(msg string, args ...any)
	// TraceContext logs at [LevelTrace] with the given context.
	TraceContext(ctx context.Context, msg string, args ...any)
	// Debug logs at [LevelDebug].
	Debug(msg string, args ...any)
	// Debugf logs at [LevelDebug].
	// Arguments are handled in the manner of [fmt.Printf].
	Debugf(msg string, args ...any)
	// DebugContext logs at [LevelDebug] with the given context.
	DebugContext(ctx context.Context, msg string, args ...any)
	// Info logs at [LevelInfo].
	Info(msg string, args ...any)
	// Infof logs at [LevelInfo].
	// Arguments are handled in the manner of [fmt.Printf].
	Infof(msg string, args ...any)
	// InfoContext logs at [LevelInfo] with the given context.
	InfoContext(ctx context.Context, msg string, args ...any)
	// Notice logs at [LevelNotice].
	Notice(msg string, args ...any)
	// Noticef logs at [LevelNotice].
	// Arguments are handled in the manner of [fmt.Printf].
	Noticef(msg string, args ...any)
	// NoticeContext logs at [LevelNotice] with the given context.
	NoticeContext(ctx context.Context, msg string, args ...any)
	// Warn logs at [LevelWarn].
	Warn(msg string, args ...any)
	// Warnf logs at [LevelWarn].
	// Arguments are handled in the manner of [fmt.Printf].
	Warnf(msg string, args ...any)
	// WarnContext logs at [LevelWarn] with the given context.
	WarnContext(ctx context.Context, msg string, args ...any)
	// Error logs at [LevelError].
	Error(msg string, args ...any)
	// Errorf logs at [LevelError].
	// Arguments are handled in the manner of [fmt.Printf].
	Errorf(msg string, args ...any)
	// ErrorContext logs at [LevelError] with the given context.
	ErrorContext(ctx context.Context, msg string, args ...any)
	// Panic logs at [LevelPanic] and then panics with the given message.
	Panic(msg string, args ...any)
	// Panicf logs at [LevelPanic] and then panics.
	// Arguments are handled in the manner of [fmt.Printf].
	Panicf(msg string, args ...any)
	// PanicContext logs at [LevelPanic] with the given context and then panics with the given message.
	PanicContext(ctx context.Context, msg string, args ...any)
	// Fatal logs at [LevelFatal] and then calls [os.Exit](1).
	Fatal(msg string, args ...any)
	// Fatalf logs at [LevelFatal] and then calls [os.Exit](1).
	// Arguments are handled in the manner of [fmt.Printf].
	Fatalf(msg string, args ...any)
	// FatalContext logs at [LevelFatal] with the given context and then calls [os.Exit](1).
	FatalContext(ctx context.Context, msg string, args ...any)

	// With returns a Logger that has the given attributes.
	With(args ...any) Provider
	// WithGroup returns a Logger that starts a group, if name is non-empty.
	// The keys of all attributes added to the Logger will be qualified by the given
	// name. (How that qualification happens depends on the [Handler.WithGroup]
	// method of the Logger's Handler.)
	//
	// If name is empty, WithGroup returns the receiver.
	WithGroup(name string) Provider

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
	// LogAttrs is a more efficient version of [Provider.Log] that accepts only Attrs.
	LogAttrs(ctx context.Context, level Level, msg string, attrs ...slog.Attr)

	// Handler returns the [slog.Handler] that the Logger emits log records to.
	Handler() slog.Handler
	// Enabled reports whether the [Provider] emits log records at the given context and level.
	Enabled(ctx context.Context, level Level) bool

	// ToSlog returns the underlying [slog.Logger].
	ToSlog() *slog.Logger
}

// logger implements the Logger interface.
// It is a wrapper around slog.Logger.
type logger struct{ *slog.Logger }

// Debug logs at LevelDebug.
func (l *logger) Debug(msg string, a ...any) {
	l.logAttrs(context.Background(), LevelDebug, msg, a...)
}

// DebugContext logs at LevelDebug.
func (l *logger) DebugContext(ctx context.Context, msg string, a ...any) {
	l.logAttrs(ctx, LevelDebug, msg, a...)
}

// Info logs at LevelInfo.
func (l *logger) Info(msg string, a ...any) {
	l.logAttrs(context.Background(), LevelInfo, msg, a...)
}

// InfoContext logs at LevelInfo.
func (l *logger) InfoContext(ctx context.Context, msg string, a ...any) {
	l.logAttrs(ctx, LevelInfo, msg, a...)
}

// Warn logs at LevelWarn.
func (l *logger) Warn(msg string, a ...any) {
	l.logAttrs(context.Background(), LevelWarn, msg, a...)
}

// WarnContext logs at LevelWarn.
func (l *logger) WarnContext(ctx context.Context, msg string, a ...any) {
	l.logAttrs(ctx, LevelWarn, msg, a...)
}

// Error logs at LevelError.
func (l *logger) Error(msg string, a ...any) {
	l.logAttrs(context.Background(), LevelError, msg, a...)
}

// ErrorContext logs at LevelError.
func (l *logger) ErrorContext(ctx context.Context, msg string, a ...any) {
	l.logAttrs(ctx, LevelError, msg, a...)
}

// With calls Logger.With on the default logger.
func (l *logger) With(a ...any) Provider {
	return &logger{Logger: l.Logger.With(a...)}
}

// WithGroup returns a Logger that starts a group, if name is non-empty.
func (l *logger) WithGroup(name string) Provider {
	return &logger{Logger: l.Logger.WithGroup(name)}
}

// Log emits a log record with the current time and the given level and message.
func (l *logger) Log(ctx context.Context, level Level, msg string, a ...any) {
	l.Logger.Log(ctx, slog.Level(level), msg, a...)
}

// LogAttrs is a more efficient version of [Provider.Log] that accepts only Attrs.
func (l *logger) LogAttrs(ctx context.Context, level Level, msg string, attrs ...slog.Attr) {
	l.Logger.LogAttrs(ctx, slog.Level(level), msg, attrs...)
}

// Enabled reports whether the [Provider] emits log records at the given context and level.
func (l *logger) Enabled(ctx context.Context, level Level) bool {
	return l.Logger.Enabled(ctx, slog.Level(level))
}

// logAttrs emits a log record with the current time and the given level, message, and attributes.
// Must be called by a public log method to ensure that the caller is correct.
func (l *logger) logAttrs(ctx context.Context, level Level, msg string, a ...any) {
	if !l.Enabled(ctx, level) {
		return
	}

	// skip is the number of stack frames to skip to find the caller.
	// We need to skip calling runtime.Callers, this function and the public log function.
	const skip = 3
	var pcs [1]uintptr
	_ = runtime.Callers(skip, pcs[:])
	r := slog.NewRecord(time.Now(), slog.Level(level), msg, pcs[0])
	r.Add(a...)
	if ctx == nil {
		ctx = context.Background()
	}

	_ = l.Handler().Handle(ctx, r)
}
