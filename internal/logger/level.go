package logger

import (
	"log/slog"
	"strings"
)

// Level is a custom type for log levels.
type Level slog.Level

// Log levels.
const (
	LevelTrace  = Level(-8)
	LevelDebug  = Level(slog.LevelDebug)
	LevelInfo   = Level(slog.LevelInfo)
	LevelNotice = Level(slog.Level(2))
	LevelWarn   = Level(slog.LevelWarn)
	LevelError  = Level(slog.LevelError)
	LevelPanic  = Level(slog.Level(12))
	LevelFatal  = Level(slog.Level(16))
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

// newLevel returns the log level based on the provided string.
// Returns [LevelInfo] if the level is not recognized.
func newLevel(level string) Level {
	switch strings.ToUpper(level) {
	case "TRACE":
		return LevelTrace
	case "DEBUG":
		return LevelDebug
	case "INFO":
		return LevelInfo
	case "NOTICE":
		return LevelNotice
	case "WARN", "WARNING":
		return LevelWarn
	case "ERROR":
		return LevelError
	default:
		return LevelInfo
	}
}

func (l Level) String() string {
	if s, ok := LevelNames[l]; ok {
		return s
	}
	return "UNKNOWN"
}
