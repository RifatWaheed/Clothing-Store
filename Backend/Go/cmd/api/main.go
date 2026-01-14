package main

import (
	"log"

	"clothing-store-backend/internal/config"
	"clothing-store-backend/internal/db"
	"clothing-store-backend/internal/router"
)

func main() {
	// Load environment variables & app config
	cfg := config.Load()

	// Initialize database connection pool
	dbPool, err := db.NewPostgresPool(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer dbPool.Close()

	// Setup Gin router with dependencies
	r := router.SetupRouter(cfg, dbPool)

	log.Println("Starting server on port", cfg.ServerPort)

	// Start HTTP server
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
