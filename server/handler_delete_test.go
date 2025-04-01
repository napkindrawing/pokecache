package server_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"napkindrawing.com/pokecache/models"
	"napkindrawing.com/pokecache/server"
)

func TestHandlerDelete(t *testing.T) {
	t.Parallel()

	srv := server.NewServer(3)

	sporto := models.Pokemon{
		ID:        777,
		Name:      "Sporto",
		Type:      models.PokemonTypeGhost,
		Height:    63,
		Weight:    207,
		Abilities: []models.PokemonAbility{},
	}

	errAdd := srv.Cache.Add(sporto)
	require.NoError(t, errAdd)

	assert.Equal(t, 1, srv.Cache.Count(), "cache has 1 entries")

	req := httptest.NewRequest(http.MethodDelete, "http://poke.mon/id/777", nil)
	w := httptest.NewRecorder()

	mux := srv.NewServeMux()

	mux.ServeHTTP(w, req)

	resp := w.Result()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := string(bodyBytes)

	assert.Equal(t, "", body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	assert.Equal(t, 0, srv.Cache.Count(), "cache has 0 entries after delete handler")
}
