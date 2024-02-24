package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"
)

// Debugf logs at LevelDebug.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Debugf(msg string, args ...any) {
	if !l.Enabled(context.Background(), LevelDebug) {
		return
	}
	pc := getCaller(2)
	r := slog.NewRecord(time.Now(), LevelDebug, fmt.Sprintf(msg, args...), pc)
	_ = l.Handler().Handle(context.Background(), r)
}

// Infof logs at LevelInfo.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Infof(msg string, args ...any) {
	if !l.Enabled(context.Background(), LevelInfo) {
		return
	}
	pc := getCaller(2)
	r := slog.NewRecord(time.Now(), LevelInfo, fmt.Sprintf(msg, args...), pc)
	_ = l.Handler().Handle(context.Background(), r)
}

// Warnf logs at LevelWarn.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Warnf(msg string, args ...any) {
	if !l.Enabled(context.Background(), LevelWarn) {
		return
	}
	pc := getCaller(2)
	r := slog.NewRecord(time.Now(), LevelWarn, fmt.Sprintf(msg, args...), pc)
	_ = l.Handler().Handle(context.Background(), r)
}

// Errorf logs at LevelError.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Errorf(msg string, args ...any) {
	if !l.Enabled(context.Background(), LevelError) {
		return
	}
	pc := getCaller(2)
	r := slog.NewRecord(time.Now(), LevelError, fmt.Sprintf(msg, args...), pc)
	_ = l.Handler().Handle(context.Background(), r)
}

// Panic logs at [LevelPanic] and then panics.
func (l *logger) Panic(msg string, args ...any) {
	pc := getCaller(2)
	r := slog.NewRecord(time.Now(), LevelPanic, msg, pc)
	r.Add(args...)
	_ = l.Handler().Handle(context.Background(), r)
	panic(msg)
}

// Panicf logs at LevelPanic and then panics.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Panicf(msg string, args ...any) {
	fmsg := fmt.Sprintf(msg, args...)
	pc := getCaller(2)
	r := slog.NewRecord(time.Now(), LevelPanic, fmsg, pc)
	_ = l.Handler().Handle(context.Background(), r)
	panic(fmsg)
}

// PanicContext logs at [LevelPanic] with the given context and then panics.
func (l *logger) PanicContext(ctx context.Context, msg string, args ...any) {
	pc := getCaller(2)
	r := slog.NewRecord(time.Now(), LevelPanic, msg, pc)
	r.Add(args...)
	_ = l.Handler().Handle(ctx, r)
	panic(msg)
}

// Fatal logs at [LevelFatal] and then calls os.Exit(1).
func (l *logger) Fatal(msg string, args ...any) {
	pc := getCaller(2)
	r := slog.NewRecord(time.Now(), LevelFatal, msg, pc)
	r.Add(args...)
	_ = l.Handler().Handle(context.Background(), r)
	os.Exit(1)
}

// Fatalf logs at LevelFatal and then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Fatalf(msg string, args ...any) {
	pc := getCaller(2)
	r := slog.NewRecord(time.Now(), LevelPanic, fmt.Sprintf(msg, args...), pc)
	_ = l.Handler().Handle(context.Background(), r)
	os.Exit(1)
}

// FatalContext logs at [LevelFatal] with the given context and then calls os.Exit(1).
func (l *logger) FatalContext(ctx context.Context, msg string, args ...any) {
	pc := getCaller(2)
	r := slog.NewRecord(time.Now(), LevelFatal, msg, pc)
	r.Add(args...)
	_ = l.Handler().Handle(ctx, r)
	os.Exit(1)
}

// getCaller returns the program counter of the caller at a given depth.
// The depth is the number of stack frames to ascend, with 0 identifying the
// getCaller function itself, 1 identifying the caller that invoked getCaller,
// and so on.
//
// Example:
//
//	pc := getCaller(1) // Returns the program counter of the caller of the function that invoked getCaller.
//	pc := getCaller(2) // Returns the program counter of the caller of the function that invoked the function that invoked getCaller.
func getCaller(depth uint8) uintptr { //nolint: unparam
	d := int(depth) + 1

	var pcs [1]uintptr
	runtime.Callers(d, pcs[:])
	return pcs[0]
}
