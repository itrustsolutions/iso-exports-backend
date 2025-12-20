package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Setup(ctx context.Context) (*pgxpool.Pool, error) {
	// Connection string
	dsn := "user=rami.k.rayya dbname=iso-experts password=12346 host=working-ubuntu port=55333 sslmode=disable"

	// Create a connection pool
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
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
