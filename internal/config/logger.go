package config

import (
	"os"
	"go.uber.org/zap"
)

var globalLogger *zap.Logger

// InitializeLogger sets up the global logger
func InitializeLogger() error {
	var err error
	globalLogger, err = zap.NewProduction() 

	if os.Getenv("ENV") == "development" {
		globalLogger, err = zap.NewDevelopment()
	}

	if err != nil {
		return err
	}
	zap.ReplaceGlobals(globalLogger)
	return nil
}

// SyncLogger flushes any buffered log entries (should be deferred)
func SyncLogger() {
	if globalLogger != nil {
		_ = globalLogger.Sync()
	}
}

