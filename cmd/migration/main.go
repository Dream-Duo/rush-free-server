package main

import (
	"flag"
	"log"

	"rush-free-server/internal/config"
	database_migration "rush-free-server/internal/database"
	"rush-free-server/pkg/database"

	_ "github.com/lib/pq" // PostgreSQL driver
	"go.uber.org/zap"
)

func main() {
	// Parse command line flags
	var (
		command = flag.String("command", "up", "migration command (up/down/version)")
		env     = flag.String("env", "development", "environment (development/staging/production)")
		version = flag.Int("version", 0, "target migration version")
	)
	flag.Parse()

	// Initialize the logger
	if err := config.InitializeLogger(); err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer config.SyncLogger() // Ensure the logger is flushed before exiting

	// Get the Data Source Name (DSN) for PostgreSQL connection
	DatabaseConfig, err := config.GetPostgresDSN(*env)
	if err != nil {
		zap.S().Fatal("failed to get the PostgreSQL DSN: %v", err)
	}

	// Connect to the database
	db, err := database.Connect(DatabaseConfig.DatabaseURL)
	if err != nil {
		zap.S().Fatal("failed to initialize database: %v", zap.Error(err))
	}
	defer func() {
		if err := db.Close(); err != nil {
			zap.S().Error("failed to close database connection", zap.Error(err))
		}
	}()

	// Create migrator
	migrator, err := database_migration.NewMigrator(db, database_migration.MigrationConfig{
		MigrationsPath: "/app/migrations/postgres",
		DatabaseURL:    DatabaseConfig.DatabaseURL,
	})
	if err != nil {
		zap.S().Fatal("failed to create migrator: %v", err)
	}
	defer func() {
		if err := migrator.Close(); err != nil {
			zap.S().Error("failed to close migrator", zap.Error(err))
		}
	}()

	// Execute command
	switch *command {
	case "up":
		if err := migrator.Up(); err != nil {
			zap.S().Fatal("migration up failed: %v", err)
		}
		zap.L().Info("Successfully ran all migrations")

	case "down":
		if err := migrator.Down(); err != nil {
			zap.S().Fatal("migration down failed: %v", err)
		}
		zap.L().Info("Successfully reverted all migrations")

	case "version":
		version, dirty, err := migrator.Version()
		if err != nil {
			zap.S().Fatal("failed to get version: %v", err)
		}
		zap.S().Info("Current migration version: %d (dirty: %v)", version, dirty)

	case "force":
		if err := migrator.Force(*version); err != nil {
			zap.S().Fatal("failed to force migration: %v", err)
		}
		zap.S().Info("Successfully forced migration to version %v", *version)

	default:
		zap.S().Fatal("unknown command: %v", *command)
	}
}
