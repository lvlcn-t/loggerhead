package logger

import (
	"context"
	"fmt"
	"os"
)

// Trace logs at [LevelTrace].
func (l *logger) Trace(msg string, args ...any) {
	l.logAttrs(context.Background(), LevelTrace, msg, args...)
}

// Tracef logs at [LevelTrace].
// Arguments are handled in the manner of [fmt.Printf].
func (l *logger) Tracef(msg string, args ...any) {
	l.logAttrs(context.Background(), LevelTrace, fmt.Sprintf(msg, args...))
}

// TraceContext logs at [LevelTrace] with the given context.
func (l *logger) TraceContext(ctx context.Context, msg string, args ...any) {
	l.logAttrs(ctx, LevelTrace, msg, args...)
}

// Debugf logs at [LevelDebug].
// Arguments are handled in the manner of [fmt.Printf].
func (l *logger) Debugf(msg string, args ...any) {
	l.logAttrs(context.Background(), LevelDebug, fmt.Sprintf(msg, args...))
}

// Infof logs at LevelInfo.
// Arguments are handled in the manner of [fmt.Printf].
func (l *logger) Infof(msg string, args ...any) {
	l.logAttrs(context.Background(), LevelInfo, fmt.Sprintf(msg, args...))
}

// Notice logs at [LevelNotice].
func (l *logger) Notice(msg string, args ...any) {
	l.logAttrs(context.Background(), LevelNotice, msg, args...)
}

// Noticef logs at [LevelNotice].
// Arguments are handled in the manner of [fmt.Printf].
func (l *logger) Noticef(msg string, args ...any) {
	l.logAttrs(context.Background(), LevelNotice, fmt.Sprintf(msg, args...))
}

// NoticeContext logs at [LevelNotice] with the given context.
func (l *logger) NoticeContext(ctx context.Context, msg string, args ...any) {
	l.logAttrs(ctx, LevelNotice, msg, args...)
}

// Warnf logs at LevelWarn.
// Arguments are handled in the manner of [fmt.Printf].
func (l *logger) Warnf(msg string, args ...any) {
	l.logAttrs(context.Background(), LevelWarn, fmt.Sprintf(msg, args...))
}

// Errorf logs at LevelError.
// Arguments are handled in the manner of [fmt.Printf].
func (l *logger) Errorf(msg string, args ...any) {
	l.logAttrs(context.Background(), LevelError, fmt.Sprintf(msg, args...))
}

// Panic logs at [LevelPanic] and then panics.
func (l *logger) Panic(msg string, args ...any) {
	l.logAttrs(context.Background(), LevelPanic, msg, args...)
	panic(msg)
}

// Panicf logs at LevelPanic and then panics.
// Arguments are handled in the manner of [fmt.Printf].
func (l *logger) Panicf(msg string, args ...any) {
	fmsg := fmt.Sprintf(msg, args...)
	l.logAttrs(context.Background(), LevelPanic, fmsg)
	panic(fmsg)
}

// PanicContext logs at [LevelPanic] and then panics.
func (l *logger) PanicContext(ctx context.Context, msg string, args ...any) {
	l.logAttrs(ctx, LevelPanic, msg, args...)
	panic(msg)
}

// exit is a variable for [os.Exit].
var exit = os.Exit

// Fatal logs at [LevelFatal] and then calls os.Exit(1).
func (l *logger) Fatal(msg string, args ...any) {
	l.logAttrs(context.Background(), LevelFatal, msg, args...)
	exit(1)
}

// Fatalf logs at LevelFatal and then calls os.Exit(1).
// Arguments are handled in the manner of [fmt.Printf].
func (l *logger) Fatalf(msg string, args ...any) {
	l.logAttrs(context.Background(), LevelFatal, fmt.Sprintf(msg, args...))
	exit(1)
}

// FatalContext logs at [LevelFatal] and then calls os.Exit(1).
func (l *logger) FatalContext(ctx context.Context, msg string, args ...any) {
	l.logAttrs(ctx, LevelFatal, msg, args...)
	exit(1)
}
