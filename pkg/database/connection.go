package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

func Connect(dataSourceName string, ctx context.Context) (*pgxpool.Pool, error) {
	// Define maximum retries and interval
	const maxRetries = 5
	const retryInterval = 5 * time.Second

	var pool *pgxpool.Pool
	var err error

	for attempt := range maxRetries {
		pool, err = pgxpool.New(ctx, dataSourceName)
		if err == nil {
			// Ping to verify the connection
			if err = pool.Ping(ctx); err == nil {
				zap.L().Info("Successfully connected to PostgreSQL database")
				return pool, nil
			}
		}

		// Log retry attempt and wait before retrying
		if attempt < maxRetries-1 {
			zap.S().Info("Failed to connect to database, retrying in %v... (%d/%d)", retryInterval, attempt+1, maxRetries)
			time.Sleep(retryInterval)
		} else {
			// Return an error if all retries fail
			return nil, fmt.Errorf("error connecting to database: %w", err)
		}
	}

	return nil, fmt.Errorf("exceeded maximum retry attempts to connect to database")
}

// Connect establishes a connection to the PostgreSQL database with retry logic.
// Returns a database connection if successful, otherwise returns an error.
func ConnectForMigration(dataSourceName string) (*sql.DB, error) {
	// Open a new database connection
	database, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	// Define maximum number of connection retries and interval between retries
	const maxRetries = 5
	const retryInterval = 5 * time.Second

	// Attempt to ping the database to ensure the connection is established
	for attempt := range maxRetries {
		if err := database.Ping(); err == nil {
			zap.L().Info("Successfully connected to PostgreSQL database")
			return database, nil
		}
		// Log retry attempt and wait before retrying
		if attempt < maxRetries-1 {
			zap.S().Info("Failed to connect to database, retrying in %v... (%d/%d)", retryInterval, attempt+1, maxRetries)
			time.Sleep(retryInterval)
		} else {
			// Return an error if all retry attempts fail
			return nil, fmt.Errorf("error pinging database: %w", err)
		}
	}

	// Return an error if the maximum number of retries is exceeded
	return nil, fmt.Errorf("exceeded maximum retry attempts to connect to database")
}
