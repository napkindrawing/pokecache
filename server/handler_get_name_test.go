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

func TestHandlerGetName(t *testing.T) {
	t.Parallel()

	srv := server.NewServer(3)

	spinoza := models.Pokemon{
		ID:        888,
		Name:      "Spinoza",
		Type:      models.PokemonTypePsychic,
		Height:    67,
		Weight:    213,
		Abilities: []models.PokemonAbility{},
	}

	errAdd := srv.Cache.Add(spinoza)
	require.NoError(t, errAdd)

	assert.Equal(t, 1, srv.Cache.Count(), "cache has 1 entries")

	req := httptest.NewRequest(http.MethodGet, "http://poke.mon/name/spinoza", nil)
	w := httptest.NewRecorder()

	mux := srv.NewServeMux()

	mux.ServeHTTP(w, req)

	resp := w.Result()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := string(bodyBytes)

	assert.Equal(t, "", body)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
