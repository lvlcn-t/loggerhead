package main

import (
	"context"
	"net/http"

	"github.com/lvlcn-t/loggerhead/logger"
)

func main() {
	ctx := logger.IntoContext(context.Background(), logger.NewNamedLogger("webserver"))
	s := &http.Server{
		Addr: ":8080",
	}

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		log := logger.FromContext(r.Context())
		log.Info("Hello, world!")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	}

	http.Handle("/", logger.Middleware(ctx)(handler))

	s.ListenAndServe()
}
