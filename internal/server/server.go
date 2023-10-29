package server

import (
	"context"
	"diploma/authorization/internal/config"
	handler "diploma/authorization/internal/delivery/handlers/http"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer(cfg *config.Config, handler *handler.Handler) (*Server, error) {
	httpHandler, err := handler.Init(cfg)
	if err != nil {
		return nil, err
	}

	return &Server{
		server: &http.Server{
			Addr:           fmt.Sprintf(":%s", cfg.Service.Port),
			Handler:        httpHandler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   30 * time.Second,
			MaxHeaderBytes: 2 << 20, // 2MB
		},
	}, nil
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
