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
	"github.com/itrustsolutions/iso-exports-backend/utils/middleware"
)

func main() {
	config := config.GetConfigOrExist()

	ctx := context.Background()

	pool, err := application.DbSetup(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not set up database:", err)
		os.Exit(1)
	}
	defer pool.Close()

	identityModule := identity.NewModule(&identity.Config{
		DB: pool,
	})

	r := chi.NewRouter()

	r.Use(middleware.CorrelationID)
	r.Mount("/identity/users", identityModule.Routes)

	if err := http.ListenAndServe(config.Server.Port, r); err != nil {
		fmt.Fprintln(os.Stderr, "Could not set up HTTP server:", err)
		os.Exit(1)
	}
}
