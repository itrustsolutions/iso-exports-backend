package application

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/itrustsolutions/iso-exports-backend/utils/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func DbSetup(ctx context.Context) (*pgxpool.Pool, error) {
	config := config.GetConfigOrExist()

	// Connection string
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=%s",
		config.Database.User,
		config.Database.Name,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.SSLMode,
	)

	// Create a connection pool
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		if config.App.Env == "local" {
			return nil, fmt.Errorf("Failed to create connection pool: %w", err)
		}

		return nil, fmt.Errorf("Failed to create connection pool")
	}

	// Wait for the database to be ready
	for {
		err = pool.Ping(ctx)
		if err == nil {
			break
		}
		log.Printf("DB not ready: %v\n", err)
		time.Sleep(2 * time.Second)
	}

	log.Println("Database is ready")
	return pool, nil
}
