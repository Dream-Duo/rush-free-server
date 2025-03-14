package config

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	DatabaseURL string
	Environment string
}

// GetPostgresDSN constructs the PostgreSQL DSN using environment variables.
func GetPostgresDSN(environment string) (DatabaseConfig, error) {
	// Retrieve environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE") // Optional, default to "disable"

	// Use default value for SSL mode if not set
	if sslmode == "" {
		sslmode = "disable"
	}

	// Construct and return the DSN
	url := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	return DatabaseConfig{
		DatabaseURL: url,
		Environment: environment,
	}, nil
}
