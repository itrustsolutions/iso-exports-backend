package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	application "github.com/itrustsolutions/iso-exports-backend/cmd/internal"
	"github.com/itrustsolutions/iso-exports-backend/core/identity"
	"github.com/itrustsolutions/iso-exports-backend/utils/config"
	customcontext "github.com/itrustsolutions/iso-exports-backend/utils/context"
	"github.com/itrustsolutions/iso-exports-backend/utils/logger"
	"github.com/itrustsolutions/iso-exports-backend/utils/middleware"
)

func main() {
	config := config.GetConfigOrExist()

	ctx := context.Background()

	logger, err := logger.Initialize()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not set up logger:", err)
		os.Exit(1)
	}

	mainLogger := logger.With().Str("correlation", "background").Logger()

	ctx = customcontext.WithLogger(ctx, &mainLogger)

	pool, err := application.DbSetup(ctx)
	if err != nil {
		mainLogger.Error().Err(err).Msg("could not set up database")
		os.Exit(1)
	}
	defer pool.Close()

	r := chi.NewRouter()

	r.Use(middleware.CorrelationID)
	r.Use(middleware.RequestLoggingMiddleware(logger))
	r.Use(middleware.Recovery)

	identity.NewModule(&identity.Config{
		DB:       pool,
		Router:   r,
		HTTPPath: "/identity",
	})

	go func() {
		if err := http.ListenAndServe(config.Server.Port, r); err != nil {
			logger.Error().Err(err).Msg("could not set up HTTP server")
			os.Exit(1)
		}
	}()

	mainLogger.Info().Msg("server is running on port " + config.Server.Port)

	select {}
}
