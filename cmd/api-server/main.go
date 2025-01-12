package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"log"

	"rush-free-server/internal/config"
	"rush-free-server/cmd/api-server/handlers"
	"rush-free-server/internal/repository/postgres"
	database_initializer "rush-free-server/internal/database"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	// Initialize the logger
	if err := config.InitializeLogger(); err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer config.SyncLogger() // Ensure the logger is flushed before exiting

	// Get the Data Source Name (DSN) for PostgreSQL connection
	DatabaseConfig, err := config.GetPostgresDSN(os.Getenv("ENV"))
	if err != nil {
		zap.S().Fatal("failed to get the PostgreSQL DSN", zap.Error(err))
	}

	// Initialize database connection with migration verification
    db, err := database_initializer.InitializeDatabase(DatabaseConfig)
    if err != nil {
        zap.S().Fatal("failed to initialize database", zap.Error(err))
    }
    defer db.Close()

	// Set up signal handling for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Initialize repository and handler
	userRepo := postgres.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo, zap.S())

	// Initialize router
	router := mux.NewRouter()
	router.HandleFunc("/users", userHandler.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{name}", userHandler.GetUsersByNameHandler).Methods("GET")
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "pong"}`))
	})

	// Create server
	server := &http.Server{
		Addr:         ":8080", // Hard-coded port for simplicity
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	// Graceful shutdown
	go func() {
		<-stop

		// Received shutdown signal
		zap.L().Info("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			// Using Error instead of Fatal for graceful shutdown
			zap.S().Error("server shutdown failed", zap.Error(err))
		}
	}()

	// Start server
	zap.L().Info("Starting server on :8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		zap.S().Fatal("server failed to start", zap.Error(err))
	}

	zap.L().Info("Server stopped")
}
