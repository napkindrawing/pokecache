package server

import "net/http"

func (srv *Server) NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /id/{id}", http.HandlerFunc(srv.HandlerGetID))
	mux.Handle("POST /add", http.HandlerFunc(srv.HandlerPostAdd))
	mux.Handle("DELETE /id/{id}", http.HandlerFunc(srv.HandlerDelete))

	return mux
}
