package database

import (
	"database/sql"
	"fmt"
	"log"

	"rush-free-server/internal/config"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func Connect() (*sql.DB, error) {
	dsn := config.GetPostgresDSN()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to PostgreSQL database")
	return db, nil
}
