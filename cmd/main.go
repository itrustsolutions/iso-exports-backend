package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	application "github.com/itrustsolutions/iso-exports-backend/cmd/internal"
	"github.com/itrustsolutions/iso-exports-backend/core/identity"
	identitydtos "github.com/itrustsolutions/iso-exports-backend/core/identity/pkg/dtos"
	"github.com/itrustsolutions/iso-exports-backend/utils/config"
	customcontext "github.com/itrustsolutions/iso-exports-backend/utils/context"
	"github.com/itrustsolutions/iso-exports-backend/utils/db"
	"github.com/itrustsolutions/iso-exports-backend/utils/logger"
)

func main() {
	cfg := config.GetConfigOrExist()
	ctx := context.Background()

	log, err := logger.Initialize()
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not set up logger:", err)
		os.Exit(1)
	}
	mainLogger := log.With().Str("correlation", "background").Logger()
	ctx = customcontext.WithLogger(ctx, &mainLogger)

	pool, err := application.DbSetup(ctx)
	if err != nil {
		mainLogger.Error().Err(err).Msg("could not set up database")
		os.Exit(1)
	}
	defer pool.Close()

	r := application.NewRouter(log)

	identityModule := identity.NewModule(&identity.Config{
		DB:       pool,
		Router:   r,
		HTTPPath: "/identity",
	})

	user, err := db.ExecWithinTx(ctx, pool, func(txCtx context.Context) (*identitydtos.CreateUserResult, error) {
		return identityModule.Users.CreateUser(ctx, &identitydtos.CreateUserInput{
			Username:        "test",
			Email:           "test@example.com",
			PlainPassword:   "password",
			IsActive:        true,
			HasSystemAccess: true,
		})
	})
	fmt.Printf("user: %v\n", user)

	server := application.NewHTTPServer(cfg.Server.Port, r)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.Start(); err != nil && err.Error() != "http: server closed" {
			mainLogger.Error().Err(err).Msg("could not set up http server")
			os.Exit(1)
		}
	}()
	mainLogger.Info().Msg("http server is running on port " + cfg.Server.Port)

	<-stop
	mainLogger.Info().Msg("shutdown signal received")

	const shutdownTimeout = 30 * time.Second
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	mainLogger.Info().Msg("shutting down http server gracefully...")
	if err := server.Shutdown(shutdownCtx); err != nil {
		mainLogger.Error().Err(err).Msg("http server forced to shutdown")
	} else {
		mainLogger.Info().Msg("http server shutdown complete")
	}

	mainLogger.Info().Msg("closing database pool...")
	pool.Close()
	mainLogger.Info().Msg("database pool closed")
}
