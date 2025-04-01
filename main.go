package main

import (
	"log/slog"
	"net/http"
	"time"

	"napkindrawing.com/pokecache/server"
)

const (
	ServerReadHeaderTimeout = 5 * time.Second
	ServerReadTimeout       = 5 * time.Second
)

func main() {
	slog.Info("starting application")

	mux := server.NewServeMux()

	//nolint:exhaustruct // Going with defaults here
	srv := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: ServerReadHeaderTimeout,
		ReadTimeout:       ServerReadTimeout,
		Handler:           mux,
	}

	slog.Info("starting server")

	err := srv.ListenAndServe()
	if err != nil {
		slog.Error("error from server", "error", err)
	}
}
