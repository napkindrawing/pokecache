package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (srv *Server) HandlerGetName(w http.ResponseWriter, req *http.Request) {
	name := req.PathValue("name")

	pokemon := srv.Cache.GetByName(name)

	if pokemon == nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	pokemonJSON, errJSON := json.Marshal(pokemon)
	if errJSON != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, errWrite := w.Write(pokemonJSON)
	if errWrite != nil {
		slog.Error("error writing response", "error", errWrite)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
