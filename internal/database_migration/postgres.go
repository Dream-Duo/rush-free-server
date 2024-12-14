package database_migration

import (
	"database/sql"
	"fmt"

	"rush-free-server/internal/config"
)

// InitializeDatabase sets up the database connection and verifies migrations
func InitializeDatabase(config config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Verify migrations are up to date in non-development environments
	if config.Environment != "development" {
		migrator, err := NewMigrator(db, MigrationConfig{
			MigrationsPath: "/app/migrations/postgres",
			DatabaseURL:    config.DatabaseURL,
		})
		if err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to create migrator: %w", err)
		}
		defer migrator.Close()

		version, dirty, err := migrator.Version()
		if err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to get migration version: %w", err)
		}

		if dirty {
			db.Close()
			return nil, fmt.Errorf("database has dirty migration state at version %d", version)
		}
	}

	return db, nil
}
