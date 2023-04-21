package server

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type Server struct {
	HTTP   *http.Server
	logger *zap.SugaredLogger
	router *RouterConfig
}

func NewServer(logger *zap.SugaredLogger, router *RouterConfig) *Server {
	server := &Server{
		HTTP: &http.Server{
			Addr: ":7891",
		},
		logger: logger,
		router: router,
	}
	fmt.Println("NewServer End")
	return server
}

func (s *Server) Serve() {
	fmt.Println("Serve Method")

	s.HTTP.Handler = s.router.addRoutes()
	port := s.HTTP.Addr

	fmt.Println("Serve Method")
	s.logger.Infow("Starting http Server ", "Port", port)
	err := s.HTTP.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		s.logger.Error("Server not able to startup with error: ", err)
	}
}

// func (s *Server) Serve() {
// 	s.HTTP.Handler = s.router.addRoutes()
// 	err := s.HTTP.ListenAndServe()

// 	fmt.Println(err)
// }
