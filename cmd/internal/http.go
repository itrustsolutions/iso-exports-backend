package application

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HTTPServer struct {
	Server *http.Server
}

func NewHTTPServer(addr string, handler *chi.Mux) *HTTPServer {
	return &HTTPServer{
		Server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (s *HTTPServer) Start() error {
	return s.Server.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
