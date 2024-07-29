package logger_test

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"

	clog "github.com/charmbracelet/log"
	"github.com/lvlcn-t/loggerhead/logger"
)

func ExampleLogger_Debug() {
	var buf bytes.Buffer
	log := logger.NewLogger(logger.Options{Handler: clog.NewWithOptions(&buf, clog.Options{Level: clog.DebugLevel})})

	log.Debug("This is a debug message!")
	log.Debugf("This is a %s message!", "debug")
	log.DebugContext(context.Background(), "This is a debug message!")

	fmt.Println(buf.String())
	// Output:
	// DEBU This is a debug message!
	// DEBU This is a debug message!
	// DEBU This is a debug message!
}

func ExampleLogger_Info() {
	var buf bytes.Buffer
	log := logger.NewLogger(logger.Options{Handler: clog.NewWithOptions(&buf, clog.Options{Level: clog.DebugLevel})})

	log.Info("Hello, world!")
	log.Infof("Hello, %s!", "world")
	log.InfoContext(context.Background(), "Hello, world!")

	fmt.Println(buf.String())
	// Output:
	// INFO Hello, world!
	// INFO Hello, world!
	// INFO Hello, world!
}

func ExampleLogger_Warn() {
	var buf bytes.Buffer
	log := logger.NewLogger(logger.Options{Handler: clog.NewWithOptions(&buf, clog.Options{Level: clog.DebugLevel})})

	log.Warn("This is a warning!")
	log.Warnf("This is a %s!", "warning")
	log.WarnContext(context.Background(), "This is a warning!")

	fmt.Println(buf.String())
	// Output:
	// WARN This is a warning!
	// WARN This is a warning!
	// WARN This is a warning!
}

func ExampleLogger_Error() {
	var buf bytes.Buffer
	log := logger.NewLogger(logger.Options{Handler: clog.NewWithOptions(&buf, clog.Options{Level: clog.DebugLevel})})

	log.Error("This is an error!")
	log.Errorf("This is an %s!", "error")
	log.ErrorContext(context.Background(), "This is an error!")

	fmt.Println(buf.String())
	// Output:
	// ERRO This is an error!
	// ERRO This is an error!
	// ERRO This is an error!
}

func ExampleNewLogger() {
	// Create a new logger with the default configuration.
	log := logger.NewLogger()
	log.Info("Hello, world!")
}

func ExampleNewNamedLogger() {
	// Create a new logger with a name.
	log := logger.NewNamedLogger("my-logger")
	log.Info("Hello, world!")
}

func ExampleIntoContext() {
	// Create a new logger with the default configuration.
	log := logger.NewLogger()
	// Inject the logger into the context.
	_ = logger.IntoContext(context.Background(), log)
	log.Info("Hello, world!")
}

func ExampleFromContext() {
	// Create a new logger with the default configuration.
	log := logger.NewLogger()
	// Inject the logger into the context.
	ctx := logger.IntoContext(context.Background(), log)
	// Retrieve the logger from the context.
	logger.FromContext(ctx).Info("Hello, world!")
}

func ExampleFromSlog() {
	// Create a new slog.Logger.
	sl := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// Convert the slog.Logger to a logger.Logger.
	log := logger.FromSlog(sl)
	log.Info("Hello, world!")
}
