package logger_test

import (
	"context"
	"testing"

	"github.com/lvlcn-t/halog/pkg/logger"
)

func TestSomeLogging(t *testing.T) {
	log := logger.NewNamedLogger("Main")
	log.Info("MainTest")
	log = logger.NewLogger()
	log.Info("NoNamedLoggerTest")
	log.With("testKey", "testValue").Info("WithTest")
}

func TestIntoContext(t *testing.T) {
	// get named logger
	log := logger.NewNamedLogger("Test")

	// propagate logger to context
	ctx := logger.IntoContext(context.Background(), log)

	// check if logger is available in context
	if l := logger.FromContext(ctx); l != log {
		t.Errorf("Logger not found in context")
	}
}

func TestFromContext(t *testing.T) {
	// get named logger
	log := logger.NewNamedLogger("Test")

	// propagate logger to context
	ctx := logger.IntoContext(context.Background(), log)

	logFromCtx := logger.FromContext(ctx)

	logFromCtx.Info("Test")
}
