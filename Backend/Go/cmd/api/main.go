package main

import (
	"log"

	"clothing-store-backend/internal/config"
	"clothing-store-backend/internal/db"
	"clothing-store-backend/internal/router"
)

func main() {
	cfg := config.Load()

	_ = db.NewPostgresPool(cfg.DBUrl)

	r := router.SetupRouter()

	log.Println("Starting server on port", cfg.AppPort)
	r.Run(":" + cfg.AppPort)
}
