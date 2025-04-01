package server

import "net/http"

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/id/{id}", http.HandlerFunc(HandlerGetID))

	return mux
}
