package logger

import (
	"context"
	"fmt"
	"os"
)

// Debugf logs at LevelDebug.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Debugf(ctx context.Context, msg string, args ...any) {
	l.logAttrs(ctx, LevelDebug, fmt.Sprintf(msg, args...))
}

// Infof logs at LevelInfo.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Infof(ctx context.Context, msg string, args ...any) {
	l.logAttrs(ctx, LevelInfo, fmt.Sprintf(msg, args...))
}

// Warnf logs at LevelWarn.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Warnf(ctx context.Context, msg string, args ...any) {
	l.logAttrs(ctx, LevelWarn, fmt.Sprintf(msg, args...))
}

// Errorf logs at LevelError.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Errorf(ctx context.Context, msg string, args ...any) {
	l.logAttrs(ctx, LevelError, fmt.Sprintf(msg, args...))
}

// Panic logs at [LevelPanic] and then panics.
func (l *logger) Panic(ctx context.Context, msg string, args ...any) {
	l.logAttrs(ctx, LevelPanic, msg, args...)
	panic(msg)
}

// Panicf logs at LevelPanic and then panics.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Panicf(ctx context.Context, msg string, args ...any) {
	fmsg := fmt.Sprintf(msg, args...)
	l.logAttrs(ctx, LevelPanic, fmsg)
	panic(fmsg)
}

// Fatal logs at [LevelFatal] and then calls os.Exit(1).
func (l *logger) Fatal(ctx context.Context, msg string, args ...any) {
	l.logAttrs(ctx, LevelFatal, msg, args...)
	os.Exit(1)
}

// Fatalf logs at LevelFatal and then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Fatalf(ctx context.Context, msg string, args ...any) {
	l.logAttrs(ctx, LevelFatal, fmt.Sprintf(msg, args...))
	os.Exit(1)
}
