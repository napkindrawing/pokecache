package server

import "napkindrawing.com/pokecache/services/cache"

type Server struct {
	Cache cache.Service
}

func NewServer(maxCapacity int) *Server {
	cacheSvc := cache.NewService(maxCapacity)

	srv := &Server{
		Cache: cacheSvc,
	}

	return srv
}
