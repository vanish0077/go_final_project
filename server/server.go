package server

import (
	"net/http"

	"go_final_project/config"
)

type Server struct {
	httpServer *http.Server
	Handler    http.Handler
}

var port = config.Port()

func (s *Server) Run(router http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    port,
		Handler: router,
	}

	return s.httpServer.ListenAndServe()
}
