package main

import (
	"flag"
	"log"
	"fmt"

	"rush-free-server/internal/config"
	"rush-free-server/internal/database_migration"
	"rush-free-server/pkg/database"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Parse command line flags
	var (
		command = flag.String("command", "up", "migration command (up/down/version)")
		env     = flag.String("env", "development", "environment (development/staging/production)")
		version = flag.Int("version", 0, "target migration version")
	)
	flag.Parse()

	fmt.Println("Command:", *command)
	fmt.Println("Environment:", *env)
	fmt.Println("Version:", *version)

	// Get the Data Source Name (DSN) for PostgreSQL connection
	DatabaseConfig, err := config.GetPostgresDSN(*env)

	// Connect to the database
	db, err := database.Connect(DatabaseConfig.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create migrator
	migrator, err := database_migration.NewMigrator(db, database_migration.MigrationConfig{
		MigrationsPath: "/app/migrations/postgres",
		DatabaseURL:    DatabaseConfig.DatabaseURL,
	})
	if err != nil {
		log.Fatalf("Failed to create migrator: %v", err)
	}
	defer migrator.Close()

	// Execute command
	switch *command {
	case "up":
		if err := migrator.Up(); err != nil {
			log.Fatalf("Migration up failed: %v", err)
		}
		log.Println("Successfully ran all migrations")

	case "down":
		if err := migrator.Down(); err != nil {
			log.Fatalf("Migration down failed: %v", err)
		}
		log.Println("Successfully reverted all migrations")

	case "version":
		version, dirty, err := migrator.Version()
		if err != nil {
			log.Fatalf("Failed to get version: %v", err)
		}
		log.Printf("Current migration version: %d (dirty: %v)", version, dirty)
	
	case "force":
		if err:= migrator.Force(*version); err != nil {
			log.Fatalf("Failed to force migration: %v", err)
		}
		log.Printf("Successfully forced migration to version %v", *version)

	default:
		log.Fatalf("Unknown command: %s", *command)
	}
}
