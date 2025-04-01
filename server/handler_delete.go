package server

import (
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"napkindrawing.com/pokecache/models"
)

func (srv *Server) HandlerDelete(w http.ResponseWriter, req *http.Request) {
	idStr := req.PathValue("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	errDel := srv.Cache.Delete(models.PokemonID(id))
	if errDel != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := io.WriteString(w, "error deleting pokemon\n")
		if errWrite != nil {
			slog.Error("error writing error body in response", "error", errWrite)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
}
