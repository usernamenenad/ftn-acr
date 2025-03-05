package server

import (
	"net/http"

	"github.com/usernamenenad/ftn-acr/router/v1"
)

type Server struct {
	server *http.Server
}

func NewServer(addr string) *Server {
	router := router.NewDefaultRouter()
	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: router.Mux,
		},
	}
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}
