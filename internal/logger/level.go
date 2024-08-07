package logger

import (
	"log/slog"
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

// LevelNames is a map of log levels to their respective names.
var LevelNames = map[Level]string{
	LevelTrace:  "TRACE",
	LevelDebug:  "DEBUG",
	LevelInfo:   "INFO",
	LevelNotice: "NOTICE",
	LevelWarn:   "WARN",
	LevelError:  "ERROR",
	LevelPanic:  "PANIC",
	LevelFatal:  "FATAL",
}

// LevelColors is a map of log levels to their respective ansi color codes.
var LevelColors = map[Level]string{
	LevelTrace:  "240", // TRACE - Light Gray
	LevelDebug:  "63",  // DEBUG - Blue
	LevelInfo:   "86",  // INFO - Cyan
	LevelNotice: "220", // NOTICE - Yellow
	LevelWarn:   "192", // WARN - Orange
	LevelError:  "204", // ERROR - Red
	LevelPanic:  "134", // PANIC - Purple
	LevelFatal:  "160", // FATAL - Dark Red
}

// getLevel returns the integer value of the given level string.
// If the level is not recognized, it returns LevelInfo.
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

// getLevelString returns the string value of the given level.
func getLevelString(level Level) string {
	if s, ok := LevelNames[level]; ok {
		return s
	}
	return "UNKNOWN"
}
