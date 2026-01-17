# Copilot Instructions for Clothing Store Backend

## Architecture Overview
Go backend using **Gin** web framework + **PostgreSQL**. Implements clean architecture with dependency injection:
- **auth/**: User registration/login with JWT (HS256, context-extracted user_id)
- **config/**: Environment variables via godotenv (DATABASE_URL, JWT_SECRET, SERVER_PORT)
- **db/**: PostgreSQL pgx/v5 connection pool (initialized once in main)
- **middleware/**: JWT validation with Bearer token parsing
- **router/**: Route setup with dependency injection pattern

**Key pattern**: Handler → Service → Repository layers. Services contain business logic (bcrypt hashing, validation), repositories execute SQL via pgx pool.

## Development Setup
Required env vars (.env or shell):
- `DATABASE_URL=postgres://user:pass@localhost:5432/clothing_store`
- `JWT_SECRET=your_secret_key`
- `SERVER_PORT=8080`

Commands:
- **Migrations** (run once): `go run cmd/migrate/main.go up`
- **Dev server**: `go run cmd/api/main.go`
- **Build**: `go build -o api ./cmd/api`
- **Docker**: `docker build -t clothing-store .`

## Request/Response Patterns
Auth endpoints (no JWT required):
- `POST /api/auth/register` → `{email, password}` → `{message}`
- `POST /api/auth/login` → `{email, password}` → `{token}`

Protected routes: Apply `middleware.JWTAuth(cfg.JWTSecret)` to groups. Extract user_id via `c.GetString("user_id")`. Header format: `Authorization: Bearer <token>`

Handler template:
```go
func (h *Handler) SomeRoute(c *gin.Context) {
    userID := c.GetString("user_id")
    var req struct{ Field string }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid"})
        return
    }
}
```

## Database & Models
Schema (migrations/000001_init.up.sql):
- `users`: id UUID, email UNIQUE, password_hash, created_at
- `products`: id UUID, name, price, created_at

Query pattern in Repository:
```go
row := r.DB.QueryRow(ctx, "SELECT id, email FROM users WHERE email = $1", email)
err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)
```

Services return errors; handlers translate to JSON responses with appropriate HTTP status codes.

## Code Conventions
- **Constructors**: `NewXxx(deps) *Xxx` (e.g., `NewHandler(service)`)
- **Context**: Pass `c.Request.Context()` to services/repos in handlers
- **Imports**: stdlib → third-party → internal
- **Error handling**: All DB calls use context.Context; errors propagate up through Service layer
- **Module name**: `clothing-store-backend`

## Key Files
- `cmd/api/main.go`: Entry point, dependency wiring
- `internal/router/router.go`: Route groups with auth middleware
- `internal/auth/service.go`: Business logic (bcrypt, user lookup)
- `migrations/000001_init.up.sql`: Schema with users + products</content>
<parameter name="filePath">d:\Clothing Store\Clothing-Store\Backend\Go\.github\copilot-instructions.md