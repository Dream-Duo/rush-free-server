package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"rush-free-server/internal/config"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Connect establishes a connection to the PostgreSQL database with retry logic.
// Returns a database connection if successful, otherwise returns an error.
func Connect() (*sql.DB, error) {
	// Get the Data Source Name (DSN) for PostgreSQL connection
	dataSourceName := config.GetPostgresDSN()

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
			log.Println("Successfully connected to PostgreSQL database")
			return database, nil
		}
		// Log retry attempt and wait before retrying
		if attempt < maxRetries-1 {
			log.Printf("Failed to connect to database, retrying in %v... (%d/%d)", retryInterval, attempt+1, maxRetries)
			time.Sleep(retryInterval)
		} else {
			// Return an error if all retry attempts fail
			return nil, fmt.Errorf("error pinging database: %w", err)
		}
	}

	// Return an error if the maximum number of retries is exceeded
	return nil, fmt.Errorf("exceeded maximum retry attempts to connect to database")
}
