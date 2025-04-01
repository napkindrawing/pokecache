package server

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"napkindrawing.com/pokecache/models"
)

func (srv *Server) HandlerPostAdd(w http.ResponseWriter, req *http.Request) {
	var pokemon models.Pokemon

	errDecode := json.NewDecoder(req.Body).Decode(&pokemon)
	if errDecode != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("error decoding request", "error", errDecode)

		return
	}

	slog.Info("Adding Pokemon", "pokemon", pokemon)

	errAdd := srv.Cache.Add(pokemon)
	if errAdd != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := io.WriteString(w, "error adding pokemon\n")
		if errWrite != nil {
			slog.Error("error writing error body in response", "error", errWrite)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
}
