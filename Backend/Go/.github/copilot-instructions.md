# Copilot Instructions for Clothing Store Backend

## Architecture Overview
This is a Go backend using Gin web framework with PostgreSQL. Follows clean architecture with internal modules:
- `auth/`: Handles user registration/login with JWT tokens
- `config/`: Loads environment variables (DATABASE_URL, JWT_SECRET, SERVER_PORT)
- `db/`: PostgreSQL connection pool using pgx/v5
- `middleware/`: JWT authentication middleware
- `router/`: Route setup with dependency injection

Key pattern: Each module has Handler/Service/Repository layers. Services handle business logic, Repositories handle DB queries.

## Dependencies & Setup
- **Database**: PostgreSQL with migrations in `migrations/` using golang-migrate
- **Auth**: JWT tokens (HS256, 24h expiry) with bcrypt password hashing
- **Config**: Environment variables via godotenv (.env file optional)

Run migrations: `go run cmd/migrate/main.go up`

## Build & Run Commands
- **Dev server**: `go run cmd/api/main.go`
- **Build binary**: `go build -o api ./cmd/api`
- **Docker**: `docker build -t clothing-store .` (exposes port 8080)

Server starts on port from SERVER_PORT env var.

## Authentication Patterns
- **Register/Login**: POST /auth/register, /auth/login with JSON {email, password}
- **Protected routes**: Use `middleware.JWTAuth(secret)` on route groups
- **Token format**: Bearer <token> in Authorization header
- **User context**: Middleware sets `user_id` in Gin context

Example handler:
```go
func (h *Handler) SomeProtected(c *gin.Context) {
    userID := c.GetString("user_id")  // From middleware
    // ... logic
}
```

## Database Patterns
- Use pgx pool for queries
- Repository methods take context.Context
- User model: {ID uuid, Email, PasswordHash, CreatedAt}

Example query:
```go
row := r.DB.QueryRow(ctx, "SELECT id, email FROM users WHERE id = $1", userID)
```

## Code Conventions
- Struct constructors: `NewXxx(deps) *Xxx`
- Error handling: Return errors from services, handlers respond with JSON
- Imports: Standard library first, then third-party, then internal
- Module name: `clothing-store-backend`

## Key Files
- `cmd/api/main.go`: Entry point, wires dependencies
- `internal/router/router.go`: Route setup with auth injection
- `internal/auth/jwt.go`: Token generation/validation
- `migrations/000001_init.up.sql`: Users and products tables</content>
<parameter name="filePath">d:\Clothing Store\Clothing-Store\Backend\Go\.github\copilot-instructions.md