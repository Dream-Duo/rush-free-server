package database_migration

import (
	"database/sql"
	"fmt"

	"rush-free-server/internal/config"
	"rush-free-server/pkg/database"

	"go.uber.org/zap"
)

// InitializeDatabase sets up the database connection and verifies migrations
func InitializeDatabase(config config.DatabaseConfig) (*sql.DB, error) {
	// Connect to the database
	db, err := database.Connect(config.DatabaseURL)
	if err != nil {
		// Return the error to the caller and allow the caller to decide whether to retry, log, or exit the program.
		zap.S().Error("failed to initialize database: %v", zap.Error(err))
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}
	defer db.Close()

	// Verify migrations are up to date in non-development environments
	if config.Environment != "development" {
		migrator, err := NewMigrator(db, MigrationConfig{
			MigrationsPath: "/app/migrations/postgres",
			DatabaseURL:    config.DatabaseURL,
		})
		if err != nil {
			db.Close()
			// Return the error to the caller
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
			return nil, fmt.Errorf("database has dirty migration state at version: %d", version)
		}
	}

	return db, nil
}
