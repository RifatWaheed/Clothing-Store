package product

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

// GetProductsPublic handles GET /api/public/products (NO AUTH REQUIRED)
// Safe for public access - read-only, paginated, rate-limited via query params
func (h *Handler) GetProductsPublic(c *gin.Context) {
	var params ListParams

	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	search := c.DefaultQuery("search", "")

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 || limitInt > 100 {
		limitInt = 20
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil || offsetInt < 0 {
		offsetInt = 0
	}

	params.Limit = limitInt
	params.Offset = offsetInt
	params.Search = search

	result, err := h.service.ListProducts(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetProductByIDPublic handles GET /api/public/products/:id (NO AUTH REQUIRED)
func (h *Handler) GetProductByIDPublic(c *gin.Context) {
	pkidStr := c.Param("id")

	pkid, err := strconv.ParseInt(pkidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	product, err := h.service.GetProduct(c.Request.Context(), pkid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetProducts handles GET /api/products (AUTH REQUIRED)
func (h *Handler) GetProducts(c *gin.Context) {
	// Parse query parameters with validation
	var params ListParams

	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	search := c.DefaultQuery("search", "")

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 || limitInt > 100 {
		limitInt = 20
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil || offsetInt < 0 {
		offsetInt = 0
	}

	params.Limit = limitInt
	params.Offset = offsetInt
	params.Search = search

	result, err := h.service.ListProducts(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetProductByID handles GET /api/products/:id (AUTH REQUIRED)
func (h *Handler) GetProductByID(c *gin.Context) {
	pkidStr := c.Param("id")

	pkid, err := strconv.ParseInt(pkidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	product, err := h.service.GetProduct(c.Request.Context(), pkid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch product"})
		return
	}

	c.JSON(http.StatusOK, product)
}
