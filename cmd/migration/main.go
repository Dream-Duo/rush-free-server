package main

import (
	"database/sql"
	"log"

	"rush-free-server/internal/config" // Adjust import path

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Get the PostgreSQL DSN from the config
	dsn := config.GetPostgresDSN()

	// Connect to the database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Ping the database to ensure the connection is successful
	if err := db.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	// Create the User table
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create the User table: %v", err)
	}

	log.Println("User table created successfully!")
}
