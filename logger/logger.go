package logger

import (
	"github.com/lvlcn-t/loggerhead/internal/logger"
)

type Logger = logger.Logger

var NewLogger = logger.NewLogger

var NewNamedLogger = logger.NewNamedLogger

var NewContextWithLogger = logger.NewContextWithLogger

var IntoContext = logger.IntoContext

var FromContext = logger.FromContext

var Middleware = logger.Middleware
