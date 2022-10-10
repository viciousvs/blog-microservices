package http

import (
	"context"
	"github.com/viciousvs/blog-microservices/gateway/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg config.ServerConfig, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           cfg.Addr,
		Handler:        handler,
		MaxHeaderBytes: cfg.MaxHeaderbytes,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
