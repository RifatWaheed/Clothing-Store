package main

import (
	"context"
	"log"

	"clothing-store-backend/internal/config"
	"clothing-store-backend/internal/db"
	"clothing-store-backend/internal/router"
)

func main() {
	cfg := config.Load()

	dbPool, err := db.NewPostgresPool(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer dbPool.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := router.SetupRouter(ctx, cfg, dbPool)

	log.Println("Starting server on port", cfg.ServerPort)

	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
