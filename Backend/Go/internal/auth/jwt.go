package auth

import (
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func init() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET is not set")
	}
	jwtSecret = []byte(secret)
}

func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

type RateLimiter struct {
	requests map[string][]time.Time
	mu       sync.Mutex
	limit    int
	window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		rl.mu.Lock()
		now := time.Now()

		// Clean old requests
		if requests, exists := rl.requests[ip]; exists {
			var valid []time.Time
			for _, t := range requests {
				if now.Sub(t) < rl.window {
					valid = append(valid, t)
				}
			}
			rl.requests[ip] = valid
		}

		// Check limit
		if len(rl.requests[ip]) >= rl.limit {
			rl.mu.Unlock()
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			c.Abort()
			return
		}

		// Add request
		rl.requests[ip] = append(rl.requests[ip], now)
		rl.mu.Unlock()

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	rateLimiter := NewRateLimiter(100, time.Minute) // 100 req/min
	r.Use(rateLimiter.Middleware())

	return r
}
