package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/lvlcn-t/loggerhead/logger"
)

type Logger interface {
	logger.Logger
	Success(msg string, args ...any)
}

const LevelSuccess = slog.Level(1)

type loggerExtension struct {
	logger.Logger
}

// Success logs at LevelSuccess.
//
// Note: The "Success" level used in this code is not part of RFC5424, which defines the standard syslog message format.
// Therefore, it is not recommended to implement a "Success" level in production systems that adhere strictly to RFC5424.
// This is just an example of how to extend the logger with custom levels.
// For more information, see https://datatracker.ietf.org/doc/html/rfc5424.
func (l *loggerExtension) Success(msg string, args ...any) {
	l.Log(context.Background(), LevelSuccess, msg, args...)
}

func NewLogger() Logger {
	return &loggerExtension{
		Logger: logger.NewLogger(logger.Opts{
			Handler: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				ReplaceAttr: replaceAttr,
			}),
		}),
	}
}

// replaceAttr replaces the value of the Level attribute with a string value.
// This is used to replace the Level attribute with a string value for the LevelSuccess level.
// If we don't do this, the Level attribute would be displayed as the closest default slog level,
// plus an offset (e.g. "INFO+1").
func replaceAttr(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		l := a.Value.Any().(slog.Level)
		if l == LevelSuccess {
			a.Value = slog.StringValue("SUCCESS")
		}
	}
	return a
}

func main() {
	log := NewLogger()
	log.Success("Hello, world!")
}
