package config

import (
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

var globalLogger *zap.Logger

// InitializeLogger loads the logger configuration from an environment-specific YAML file
func InitializeLogger() error {
	// Determine the environment (default to production)
	env := os.Getenv("ENV")
	configFilePath := os.Getenv("FILE_PATH")
  if env == "" {
		env = "production"
	}

	// Read the configuration file
	file, err := os.Open(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to open logger config file: %w", err)
	}
	defer file.Close()

	// Decode the configuration into a zap.Config struct
	var cfg zap.Config
	if filepath.Ext(configFilePath) == ".yaml" || filepath.Ext(configFilePath) == ".yml" {
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&cfg); err != nil {
			return fmt.Errorf("failed to parse logger config YAML: %w", err)
		}
	} else {
		return fmt.Errorf("unsupported file format: %s", configFilePath)
	}

	// Build the logger from the configuration
	globalLogger, err = cfg.Build()
	if err != nil {
		return fmt.Errorf("failed to build logger from config: %w", err)
	}

	// Replace global logger
	zap.ReplaceGlobals(globalLogger)
	return nil
}

// SyncLogger flushes any buffered log entries
func SyncLogger() {
	if globalLogger != nil {
		_ = globalLogger.Sync()
	}
}
