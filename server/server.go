package server

import (
	"go_final_project/config"
	"net/http"
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
