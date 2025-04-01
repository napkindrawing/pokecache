package server

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

func HandlerGetID(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	_, err := io.WriteString(w, fmt.Sprintf("Hello, '%s'!\n", id))
	if err != nil {
		slog.Error("error writing response", "error", err)
	}
}
