package logger

import (
	"github.com/lvlcn-t/loggerhead/internal/logger"
)

// Logger is the interface for the logger.
// Its build on top of slog.Logger and extends it with additional logging methods.
type Logger = logger.Logger

// NewLogger creates a new Logger instance.
// If handlers are provided, the first handler in the slice is used; otherwise,
// a default JSON handler writing to os.Stderr is used. This function allows for
// custom configuration of logging handlers.
//
// Example:
//
//	log := logger.NewLogger()
//	log.Info("Hello, world!")
var NewLogger = logger.NewLogger

// NewNamedLogger creates a new Logger instance with the provided name.
// If handlers are provided, the first handler in the slice is used; otherwise,
// a default JSON handler writing to os.Stderr is used. This function allows for
// custom configuration of logging handlers.
var NewNamedLogger = logger.NewNamedLogger

// NewContextWithLogger creates a new context based on the provided parent context.
// It embeds a logger into this new context, which is a child of the logger from the parent context.
// The child logger inherits settings from the parent.
// Returns the child context and its cancel function to cancel the new context.
var NewContextWithLogger = logger.NewContextWithLogger

// IntoContext embeds the provided slog.Logger into the given context and returns the modified context.
// This function is used for passing loggers through context, allowing for context-aware logging.
var IntoContext = logger.IntoContext

// FromContext extracts the slog.Logger from the provided context.
// If the context does not have a logger, it returns a new logger with the default configuration.
// This function is useful for retrieving loggers from context in different parts of an application.
var FromContext = logger.FromContext

// Middleware takes the logger from the context and adds it to the request context.
var Middleware = logger.Middleware
