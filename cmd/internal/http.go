package application

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/itrustsolutions/iso-exports-backend/utils/config"
)

func HTTPSetup() (*chi.Mux, error) {
	config := config.GetConfigOrExist()

	r := chi.NewRouter()

	err := http.ListenAndServe(config.Server.Port, r)

	if err != nil {
		return nil, fmt.Errorf("HTTP server failed to start: %w", err)
	}

	return r, nil
}
