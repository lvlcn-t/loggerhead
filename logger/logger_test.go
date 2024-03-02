package logger_test

import (
	"context"
	"testing"

	"github.com/lvlcn-t/loggerhead/logger"
)

func TestSomeLogging(_ *testing.T) {
	log := logger.NewNamedLogger("Main")
	log.Info("MainTest")
	log = logger.NewLogger()
	log.Info("NoNamedLoggerTest")
	log.With("testKey", "testValue").Info("WithTest")
}

func TestIntoContext(t *testing.T) {
	log := logger.NewNamedLogger("Test")
	ctx := logger.IntoContext(context.Background(), log)
	if l := logger.FromContext(ctx); l != log {
		t.Errorf("Logger not found in context")
	}
}

func TestFromContext(_ *testing.T) {
	log := logger.NewNamedLogger("Test")
	ctx := logger.IntoContext(context.Background(), log)
	loc := logger.FromContext(ctx)
	loc.Info("Test")
}

func TestSomeLoggingWithEnv(t *testing.T) {
	t.Setenv("LOG_LEVEL", "debug")
	t.Setenv("LOG_FORMAT", "text")
	log := logger.NewLogger()
	log.Info("Test")
}
