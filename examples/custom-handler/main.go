package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/lvlcn-t/loggerhead/logger"
)

func main() {
	// Setups a custom handler for the logger.
	clog := log.New(os.Stdout)
	clog.SetLevel(log.DebugLevel)

	// Create a new logger with the custom handler.
	l := logger.NewLogger(clog)

	// Log some messages.
	l.Debug("I'm not sure what's happening.")
	l.Info("Hello, world in pretty colors!")
	l.Warn("I'm warning you!")
	l.Error("I'm sorry. I'm afraid you did something wrong.")
	func() {
		defer func() {
			if r := recover(); r != nil {
				l.Fatal("I'm dying!")
			}
		}()
		l.Panic("I'm panicking!")
	}()
}
