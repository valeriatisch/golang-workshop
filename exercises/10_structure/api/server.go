package api

import s "github.com/valeriatisch/golang-workshop/exercises/10_structure/storage"

// Server represents the API server
type Server struct {
	storage s.Storage
}

// NewServer creates a new instance of the Server
func NewServer(storage s.Storage) *Server {
	return &Server{
		storage: storage,
	}
}
