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
	// Initialize logger
	zapLogger, _ := zap.NewDevelopment() // Development logger for better readability
	defer zapLogger.Sync()
	logger := zapLogger.Sugar()

	// Connect to the database
	db, err := database.Connect(config.DatabaseURL)
	if err != nil {
		logger.Fatal("Failed to initialize database: %v", err)
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
