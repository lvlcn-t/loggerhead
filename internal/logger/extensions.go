package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type Level = slog.Level

const (
	LevelTrace   Level = slog.Level(-8)
	LevelDebug   Level = slog.LevelDebug
	LevelInfo    Level = slog.LevelInfo
	LevelNotice  Level = slog.Level(2)
	LevelWarning Level = slog.LevelWarn
	LevelError   Level = slog.LevelError
	LevelPanic   Level = slog.Level(12)
	LevelFatal   Level = slog.Level(16)
)

func (l *logger) Debugf(msg string, args ...any) {
	formattedMsg := fmt.Sprintf(msg, args...)
	l.Debug(formattedMsg)
}

func (l *logger) Infof(msg string, args ...any) {
	formattedMsg := fmt.Sprintf(msg, args...)
	l.Info(formattedMsg)
}

func (l *logger) Warnf(msg string, args ...any) {
	formattedMsg := fmt.Sprintf(msg, args...)
	l.Warn(formattedMsg)
}

func (l *logger) Errorf(msg string, args ...any) {
	formattedMsg := fmt.Sprintf(msg, args...)
	l.Error(formattedMsg)
}

func (l *logger) Panic(msg string, args ...any) {
	l.Log(context.Background(), LevelPanic, msg, args...)
	panic(msg)
}

func (l *logger) Panicf(msg string, args ...any) {
	formattedMsg := fmt.Sprintf(msg, args...)
	l.Log(context.Background(), LevelPanic, formattedMsg)
	panic(formattedMsg)
}

func (l *logger) PanicContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelPanic, msg, args...)
	panic(msg)
}

func (l *logger) Fatal(msg string, args ...any) {
	l.Log(context.Background(), LevelFatal, msg, args...)
	os.Exit(1)
}

func (l *logger) Fatalf(msg string, args ...any) {
	formattedMsg := fmt.Sprintf(msg, args...)
	l.Log(context.Background(), LevelFatal, formattedMsg)
	os.Exit(1)
}

func (l *logger) FatalContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelFatal, msg, args...)
	os.Exit(1)
}

// getLevel takes a level string and maps it to the corresponding Level
// Returns the level if no mapped level is found it returns info level
func getLevel(level string) int {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return int(slog.LevelDebug)
	case "INFO":
		return int(slog.LevelInfo)
	case "WARN", "WARNING":
		return int(slog.LevelWarn)
	case "ERROR":
		return int(slog.LevelError)
	default:
		return int(slog.LevelInfo)
	}
}
