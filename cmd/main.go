package main

import (
	"context"
	"fmt"
	"os"

	"github.com/itrustsolutions/iso-exports-backend/cmd/internal/db"
	"github.com/itrustsolutions/iso-exports-backend/core/identity"
)

func main() {
	ctx := context.Background()

	pool, err := db.Setup(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not set up database:", err)
		os.Exit(1)
	}
	defer pool.Close()

	identity.NewModule(&identity.Config{
		DB: pool,
	})
}
