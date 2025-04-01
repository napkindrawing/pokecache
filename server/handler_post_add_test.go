package server_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"napkindrawing.com/pokecache/server"
)

func TestHandlerPostAdd(t *testing.T) {
	t.Parallel()

	bodyReader := strings.NewReader(`{"id":12,"name":"Bulbasaur","type":"Fighting","height":123,"weight":4567,"abilities":[{"name":"Jazz Dance","description":"Dancey dance!","generation":12}]}`)

	req := httptest.NewRequest(http.MethodPost, "http://poke.mon/add", bodyReader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	srv := server.NewServer(3)

	assert.Equal(t, 0, srv.Cache.Count(), "cache has 0 entries")

	mux := srv.NewServeMux()

	mux.ServeHTTP(w, req)

	resp := w.Result()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := string(bodyBytes)

	assert.Equal(t, "", body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, 1, srv.Cache.Count(), "cache has 1 entries")
}
