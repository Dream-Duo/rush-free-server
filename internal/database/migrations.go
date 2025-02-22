package database_migration

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// MigrationConfig holds configuration for database migrations
type MigrationConfig struct {
	MigrationsPath string
	DatabaseURL    string
}

// Migrator handles database migrations
type Migrator struct {
	db       *sql.DB
	migrator *migrate.Migrate
}

// NewMigrator creates a new migrator instance
func NewMigrator(db *sql.DB, config MigrationConfig) (*Migrator, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+config.MigrationsPath,
		"postgres_db",
		driver,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrator: %w", err)
	}

	return &Migrator{
		db:       db,
		migrator: m,
	}, nil
}

// Up runs all pending migrations
func (m *Migrator) Up() error {
	err := m.migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		// Return the error to the caller
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	return nil
}

// Down rolls back all migrations
func (m *Migrator) Down() error {
	err := m.migrator.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to rollback migrations: %w", err)
	}
	return nil
}

// Version returns the current migration version
func (m *Migrator) Version() (uint, bool, error) {
	return m.migrator.Version()
}

// Close releases the migrator resources
func (m *Migrator) Close() error {
	_, err := m.migrator.Close()
	return err
}

// Force sets the migration to a specific version, even if the database is dirty
func (m *Migrator) Force(version int) error {
	err := m.migrator.Force(version)
	if err != nil {
		return fmt.Errorf("failed to force migration to version %d: %w", version, err)
	}
	return nil
}
