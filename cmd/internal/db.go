package application

import (
	"context"
	"fmt"
	"time"

	"github.com/itrustsolutions/iso-exports-backend/utils/config"
	customcontext "github.com/itrustsolutions/iso-exports-backend/utils/context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func DbSetup(ctx context.Context) (*pgxpool.Pool, error) {
	logger := customcontext.ExtractLogger(ctx)

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
		return nil, fmt.Errorf("failed to create connection pool")
	}

	var retries = 0
	// Wait for the database to be ready
	for {
		err = pool.Ping(ctx)
		if err == nil {
			break
		}

		logger.Warn().Err(err).Msg("database not ready, retrying after 2 seconds..." + fmt.Sprintf("(attempt %d)", retries+1))

		retries++
		time.Sleep(2 * time.Second)
	}

	logger.Info().Msg("database is ready")
	return pool, nil
}
