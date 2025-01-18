package database_migration

import (
	"context"
  "fmt"
  "time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// InitializeRedis sets up the Redis client connection.
func InitializeRedis() (*redis.Client, error) {
  client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", 
		Password: "",          
		DB:       0,          
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		// Return the error to the caller
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	zap.L().Info("Connected to Redis successfully")
 
  return client, nil
}
