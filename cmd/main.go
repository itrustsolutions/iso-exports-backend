package main

import (
	"context"
	"fmt"
	"os"

	application "github.com/itrustsolutions/iso-exports-backend/cmd/internal"
	"github.com/itrustsolutions/iso-exports-backend/core/identity"
	identitydtos "github.com/itrustsolutions/iso-exports-backend/core/identity/pkg/dtos"
)

func main() {
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

	user, err := identityModule.Users.CreateUser(ctx, &identitydtos.CreateUserInput{
		Username:        "test",
		Email:           "test@example.com",
		PlainPassword:   "password",
		IsActive:        true,
		HasSystemAccess: true,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not create user:", err)
		os.Exit(1)
	}

	fmt.Println("Created user:", user)
}
