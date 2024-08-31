package main

import (
	"github.com/lvlcn-t/loggerhead/logger"
)

func main() {
	// Create a new logger.
	log := logger.NewLogger(logger.Options{
		Level:  "trace",
		Format: "text",
	})

	// Log some messages.
	log.Trace("This is a trace message!")
	log.Debugf("This is a %s message!", "debug")
	log.Info("Hello, world!")
	log.Noticef("This is a %s message!", "notice")
	log.Warn("This is a warning!")
	log.Errorf("This is an %s!", "error")

	// Panic.
	func() {
		defer func() {
			if r := recover(); r != nil {
				log.Fatal("This is a fatal error!")
			}
		}()
		log.Panic("This is a panic!")
	}()
}
