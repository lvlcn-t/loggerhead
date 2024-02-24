package main

import (
	"os"

	"github.com/lvlcn-t/loggerhead/logger"
)

func main() {
	err := os.Setenv("LOG_FORMAT", "text")
	if err != nil {
		panic(err)
	}

	log := logger.NewLogger()

	log.Debug("This is a debug message!")
	log.Debugf("This is a %s message!", "debug")

	log.Info("Hello, world!")
	log.Infof("Hello, %s!", "world")

	log.Warn("This is a warning!")
	log.Warnf("This is a %s!", "warning")

	log.Error("This is an error!")
	log.Errorf("This is an %s!", "error")

	func() {
		defer func() {
			if r := recover(); r != nil {
				log.Fatal("This is a fatal error!")
			}
		}()
		log.Panic("This is a panic!")
	}()
}
