package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/itrustsolutions/iso-exports-backend/utils/middleware"
	"github.com/rs/zerolog"
)

func NewRouter(logger zerolog.Logger) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.CorrelationID)
	r.Use(middleware.RequestLoggingMiddleware(logger))
	r.Use(middleware.Recovery)
	return r
}
