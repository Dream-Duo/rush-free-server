package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
	"go.uber.org/zap"
)

// Connect establishes a connection to the PostgreSQL database with retry logic.
// Returns a database connection if successful, otherwise returns an error.
func Connect(dataSourceName string) (*sql.DB, error) {
	// Initialize logger
	zapLogger, _ := zap.NewDevelopment() // Development logger for better readability
	defer zapLogger.Sync()
	logger := zapLogger.Sugar()

	// Open a new database connection
	database, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	// Define maximum number of connection retries and interval between retries
	const maxRetries = 5
	const retryInterval = 5 * time.Second

	// Attempt to ping the database to ensure the connection is established
	for attempt := 0; attempt < maxRetries; attempt++ {
		if err := database.Ping(); err == nil {
			logger.Info("Successfully connected to PostgreSQL database")
			return database, nil
		}
		// Log retry attempt and wait before retrying
		if attempt < maxRetries-1 {
			logger.Info("Failed to connect to database, retrying in %v... (%d/%d)", retryInterval, attempt+1, maxRetries)
			time.Sleep(retryInterval)
		} else {
			// Return an error if all retry attempts fail
			return nil, fmt.Errorf("error pinging database: %w", err)
		}
	}

	// Return an error if the maximum number of retries is exceeded
	return nil, fmt.Errorf("exceeded maximum retry attempts to connect to database")
}
