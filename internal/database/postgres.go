package database_migration

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"rush-free-server/internal/config"
	"rush-free-server/pkg/database"

	"go.uber.org/zap"
)

// InitializeDatabase sets up the database connection and verifies migrations
func InitializeDatabase(config config.DatabaseConfig, ctx context.Context) (*pgxpool.Pool, error) {

	// Verify migrations are up to date in non-development environments
	if config.Environment != "development" {
		// Connect to the database
		db, err := database.ConnectForMigration(config.DatabaseURL)
		if err != nil {
			// Return the error to the caller and allow the caller to decide whether to retry, log, or exit the program.
			zap.S().Error("failed to initialize database: %v", zap.Error(err))
			return nil, fmt.Errorf("failed to initialize database: %w", err)
		}
		migrator, err := NewMigrator(db, MigrationConfig{
			MigrationsPath: "/app/migrations/database",
			DatabaseURL:    config.DatabaseURL,
		})
		if err != nil {
			if err := db.Close(); err != nil {
				zap.S().Error("failed to close database connection", zap.Error(err))
			}
			// Return the error to the caller
			return nil, fmt.Errorf("failed to create migrator: %w", err)
		}
		defer func() {
			if err := migrator.Close(); err != nil {
				zap.S().Error("failed to close migrator", zap.Error(err))
			}
		}()
		version, dirty, err := migrator.Version()
		if err != nil {
			if err := db.Close(); err != nil {
				zap.S().Error("failed to close database connection", zap.Error(err))
			}
			return nil, fmt.Errorf("failed to get migration version: %w", err)
		}

		if dirty {
			if err := db.Close(); err != nil {
				zap.S().Error("failed to close database connection", zap.Error(err))
			}
			return nil, fmt.Errorf("database has dirty migration state at version: %d", version)
		}
	}

	pool, err := database.Connect(config.DatabaseURL, ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	return pool, nil
}
