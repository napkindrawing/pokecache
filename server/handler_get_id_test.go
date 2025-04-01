package server_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"napkindrawing.com/pokecache/server"
)

func TestHandlerGetID(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "http://poke.mon/id/123", nil)
	w := httptest.NewRecorder()

	mux := server.NewServeMux()

	mux.ServeHTTP(w, req)

	resp := w.Result()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := string(bodyBytes)

	assert.Equal(t, "Hello, '123'!\n", body)
}
