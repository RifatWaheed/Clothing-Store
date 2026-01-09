package router

import (
	"clothing-store-backend/internal/config"
	"clothing-store-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	cfg := config.Load()

	authMiddleware := middleware.JWTAuth(cfg.JWTSecret)

	api := r.Group("/api")

	protected := api.Group("/")
	protected.Use(authMiddleware)
	{
		protected.GET("/me", func(c *gin.Context) {
			userID := c.GetString("user_id")
			c.JSON(200, gin.H{"user_id": userID})
		})
	}

	return r
}
