package router

import (
	"clothing-store-backend/internal/auth"
	"clothing-store-backend/internal/config"
	"clothing-store-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(cfg *config.Config, dbPool *pgxpool.Pool) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// ===== AUTH SETUP =====
	authRepo := auth.NewRepository(dbPool)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/request-sendOTP", authHandler.SendOTP)
		authRoutes.POST("/request-validateOTP", authHandler.ValidateOTP)
		authRoutes.POST("/login", authHandler.Login)
	}

	// ===== JWT PROTECTED ROUTES =====

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
