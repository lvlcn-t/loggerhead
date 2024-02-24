package main

import (
	"context"
	"net/http"
	"time"

	"github.com/lvlcn-t/loggerhead/logger"
)

func main() {
	log := logger.NewNamedLogger("webserver")
	ctx := logger.IntoContext(context.Background(), log)
	s := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 5 * time.Second,
	}

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		log := logger.FromContext(r.Context())
		log.Info("Hello, world!")

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Hello, world!"))
		if err != nil {
			log.Error("Failed to write response", "error", err)
		}
	}

	http.Handle("/", logger.Middleware(ctx)(handler))

	if err := s.ListenAndServe(); err != nil {
		log.Fatal("Failed to run server", "error", err)
	}
}
