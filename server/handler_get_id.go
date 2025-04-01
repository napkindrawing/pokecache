package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"napkindrawing.com/pokecache/models"
)

func (srv *Server) HandlerGetID(w http.ResponseWriter, req *http.Request) {
	idStr := req.PathValue("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	pokemon := srv.Cache.GetByID(models.PokemonID(id))

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
