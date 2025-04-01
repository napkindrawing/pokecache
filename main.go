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
	defaultMaxEntries       = 4
)

func main() {
	slog.Info("starting application")

	srv := server.NewServer(defaultMaxEntries)

	//nolint:exhaustruct // Going with defaults here
	httpSrv := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: ServerReadHeaderTimeout,
		ReadTimeout:       ServerReadTimeout,
		Handler:           srv.NewServeMux(),
	}

	slog.Info("starting server")

	err := httpSrv.ListenAndServe()
	if err != nil {
		slog.Error("error from server", "error", err)
	}
}
