package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

// Level is the type of log levels.
type Level = slog.Level

const (
	LevelTrace  Level = slog.Level(-8)
	LevelDebug  Level = slog.LevelDebug
	LevelInfo   Level = slog.LevelInfo
	LevelNotice Level = slog.Level(2)
	LevelWarn   Level = slog.LevelWarn
	LevelError  Level = slog.LevelError
	LevelPanic  Level = slog.Level(12)
	LevelFatal  Level = slog.Level(16)
)

// Debugf logs at LevelDebug.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Debugf(msg string, args ...any) {
	formattedMsg := fmt.Sprintf(msg, args...)
	l.Debug(formattedMsg)
}

// Infof logs at LevelInfo.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Infof(msg string, args ...any) {
	formattedMsg := fmt.Sprintf(msg, args...)
	l.Info(formattedMsg)
}

// Warnf logs at LevelWarn.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Warnf(msg string, args ...any) {
	formattedMsg := fmt.Sprintf(msg, args...)
	l.Warn(formattedMsg)
}

// Errorf logs at LevelError.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Errorf(msg string, args ...any) {
	formattedMsg := fmt.Sprintf(msg, args...)
	l.Error(formattedMsg)
}

// Panic logs at [LevelPanic] and then panics.
func (l *logger) Panic(msg string, args ...any) {
	l.Log(context.Background(), LevelPanic, msg, args...)
	panic(msg)
}

// Panicf logs at LevelPanic and then panics.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Panicf(msg string, args ...any) {
	formattedMsg := fmt.Sprintf(msg, args...)
	l.Log(context.Background(), LevelPanic, formattedMsg)
	panic(formattedMsg)
}

// PanicContext logs at [LevelPanic] with the given context and then panics.
func (l *logger) PanicContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelPanic, msg, args...)
	panic(msg)
}

// Fatal logs at [LevelFatal] and then calls os.Exit(1).
func (l *logger) Fatal(msg string, args ...any) {
	l.Log(context.Background(), LevelFatal, msg, args...)
	os.Exit(1)
}

// Fatalf logs at LevelFatal and then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Fatalf(msg string, args ...any) {
	formattedMsg := fmt.Sprintf(msg, args...)
	l.Log(context.Background(), LevelFatal, formattedMsg)
	os.Exit(1)
}

// FatalContext logs at [LevelFatal] with the given context and then calls os.Exit(1).
func (l *logger) FatalContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelFatal, msg, args...)
	os.Exit(1)
}

// getLevel takes a level string and maps it to the corresponding Level
// Returns the level if no mapped level is found it returns info level
func getLevel(level string) int {
	switch strings.ToUpper(level) {
	case "TRACE":
		return int(LevelTrace)
	case "DEBUG":
		return int(LevelDebug)
	case "INFO":
		return int(LevelInfo)
	case "NOTICE":
		return int(LevelNotice)
	case "WARN", "WARNING":
		return int(LevelWarn)
	case "ERROR":
		return int(LevelError)
	default:
		return int(LevelInfo)
	}
}
