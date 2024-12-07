package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"rush-free-server/cmd/api-server/handlers"
	"rush-free-server/internal/repository/postgres"
	"rush-free-server/pkg/database"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logger, _ := zap.NewDevelopment() // Development logger for better readability
	defer logger.Sync()
	sugar := logger.Sugar()

	db, err := database.Connect()
	if err != nil {
		logger.Fatal("Failed to connect to the database: %v", zap.Error(err))
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("Error closing the database connection: %v", zap.Error(err))
		}
	}()

	// Step 3: Set up signal handling for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Initialize repository and handler
	userRepo := postgres.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo, sugar)

	// Initialize router
	router := mux.NewRouter()
	router.HandleFunc("/users", userHandler.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{name}", userHandler.GetUsersByNameHandler).Methods("GET")

	// Basic health check route
	// curl http://localhost:8080/health
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()

	// Sample API endpoint
	// curl http://localhost:8080/api/v1/ping
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
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		// Received shutdown signal
		logger.Info("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Fatal("Server shutdown failed", zap.Error(err))
		}
	}()

	// Start server
	logger.Info("Starting server on :8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatal("Server failed to start", zap.Error(err))
	}

	logger.Info("Server stopped")
}
