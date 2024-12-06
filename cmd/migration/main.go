package main

import (
	"go.uber.org/zap"
)

// Database schema migration
func main() {
	logger, _ := zap.NewDevelopment() // Development logger for better readability
	defer logger.Sync()
	logger.Info("I am alive!")
}
