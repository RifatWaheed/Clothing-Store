package router

import (
	"clothing-store-backend/internal/auth"
	"clothing-store-backend/internal/config"
	"clothing-store-backend/internal/email"
	"clothing-store-backend/internal/middleware"
	"clothing-store-backend/internal/product"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(cfg *config.Config, dbPool *pgxpool.Pool) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// ===== EMAIL SERVICE SETUP =====
	var emailSender email.EmailSender
	if cfg.SendGridAPIKey != "" {
		emailSender = email.NewSendGridClient(cfg.SendGridAPIKey, cfg.SendGridFrom)
	} else {
		emailSender = email.NewMockEmailSender()
	}

	// ===== AUTH SETUP =====
	authRepo := auth.NewRepository(dbPool)
	authService := auth.NewService(authRepo, emailSender)
	authHandler := auth.NewHandler(authService)

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/request-sendOTP", authHandler.SendOTP)
		authRoutes.POST("/request-validateOTP", authHandler.ValidateOTP)
		authRoutes.POST("/login", authHandler.Login)
	}

	// ===== PRODUCT SETUP =====
	productRepo := product.NewRepository(dbPool)
	productService := product.NewService(productRepo)
	productHandler := product.NewHandler(productService)

	// ===== PUBLIC PRODUCT ROUTES (NO AUTH) =====
	publicRoutes := r.Group("/api/public")
	{
		publicRoutes.GET("/products", productHandler.GetProductsPublic)
		publicRoutes.GET("/products/:id", productHandler.GetProductByIDPublic)
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

		// ===== PROTECTED PRODUCT ROUTES =====
		products := protected.Group("/products")
		{
			// GET /api/products - List all products with pagination & search
			products.GET("", productHandler.GetProducts)
			// GET /api/products/:id - Get single product by ID
			products.GET("/:id", productHandler.GetProductByID)
		}
	}

	return r
}
