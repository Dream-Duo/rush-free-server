package main

import (
	"flag"

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

	// Initialize logger
	zapLogger, _ := zap.NewDevelopment() // Development logger for better readability
	defer zapLogger.Sync()
	logger := zapLogger.Sugar()

	// Get the Data Source Name (DSN) for PostgreSQL connection
	DatabaseConfig, err := config.GetPostgresDSN(*env)
	if err != nil {
		logger.Fatal("Failed to get the PostgreSQL DSN: %v", err)
	}

	// Connect to the database
	db, err := database.Connect(DatabaseConfig.DatabaseURL)
	if err != nil {
		logger.Fatal("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Create migrator
	migrator, err := database_migration.NewMigrator(db, database_migration.MigrationConfig{
		MigrationsPath: "/app/migrations/postgres",
		DatabaseURL:    DatabaseConfig.DatabaseURL,
	})
	if err != nil {
		logger.Fatal("Failed to create migrator: %v", err)
	}
	defer migrator.Close()

	// Execute command
	switch *command {
	case "up":
		if err := migrator.Up(); err != nil {
			logger.Fatal("Migration up failed: %v", err)
		}
		logger.Info("Successfully ran all migrations")

	case "down":
		if err := migrator.Down(); err != nil {
			logger.Fatal("Migration down failed: %v", err)
		}
		logger.Info("Successfully reverted all migrations")

	case "version":
		version, dirty, err := migrator.Version()
		if err != nil {
			logger.Fatal("Failed to get version: %v", err)
		}
		logger.Info("Current migration version: %d (dirty: %v)", version, dirty)
	
	case "force":
		if err:= migrator.Force(*version); err != nil {
			logger.Fatal("Failed to force migration: %v", err)
		}
		logger.Info("Successfully forced migration to version %v", *version)

	default:
		logger.Fatal("Unknown command: %v", *command)
	}
}
