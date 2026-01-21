# Copilot Instructions for Clothing Store Backend

## Architecture Overview
**Go backend** using Gin web framework + PostgreSQL with **three-layer clean architecture**:
- **Handler** (gin.Context handlers) → **Service** (business logic) → **Repository** (pgx SQL execution)
- **Dependency Injection**: Constructors pass dependencies down layers (e.g., `NewHandler(service)`, `NewService(repo)`)
- **Key modules**: `auth/` (OTP registration/login), `config/` (env vars), `middleware/` (JWT), `router/` (route setup), `db/` (pgx pool)

**Auth flow**: SendOTP (email exists check) → OTP hashed + stored (10-min expiry) → ValidateOTP (incomplete) → Login with email/password → JWT token (HS256, 24h expiry) → Protected routes via Bearer token.

## Development Setup
**Environment** (.env or export):
```
DATABASE_URL=postgres://user:pass@localhost:5432/clothing_store
JWT_SECRET=<32+ chars>
SERVER_PORT=8080
```

**Quick start**:
- `go run cmd/migrate/main.go up` — Apply migrations (run once)
- `go run cmd/api/main.go` — Start server (http://localhost:8080)
- `go build -o api ./cmd/api` — Build binary
- `docker build -t clothing-store .` — Build image

## Request/Response Patterns
**Auth endpoints** (unprotected):
- `POST /api/auth/request-sendOTP` — `{email}` → `{message}` (OTP: 6-digit alphanumeric, 10-min expiry, email must not exist)
- `POST /api/auth/request-validateOTP` — `{email, password}` → *handler incomplete*
- `POST /api/auth/login` — `{email, password}` → `{token}` (JWT valid 24h, claim: `user_id`)

**Protected routes** example:
```go
// In router.SetupRouter: apply middleware to group
protected := api.Group("/").Use(middleware.JWTAuth(cfg.JWTSecret))
protected.GET("/me", func(c *gin.Context) {
    userID := c.GetString("user_id") // Extracted by middleware
    c.JSON(200, gin.H{"user_id": userID})
})
```

**Handler pattern**: Bind JSON → validate → call service → return JSON with status code
```go
func (h *Handler) SomeRoute(c *gin.Context) {
    var req struct{ Email, Password string }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid"})
        return
    }
    result, err := h.service.SomeMethod(c.Request.Context(), req.Email)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, result)
}
```

## Database & Models
**Schema** (migrations/000001_init.up.sql):
- `users`: PK `id UUID`, `email UNIQUE`, `password_hash TEXT`, `created_at TIMESTAMP`
- `email_verifications`: `email TEXT`, `otp_hash TEXT`, `expires_at TIMESTAMP` (OTP storage)
- `products`: PK `id UUID`, `name TEXT`, `price NUMERIC(10,2)`, `created_at TIMESTAMP`

**Repository pattern** (pgx v5):
```go
func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
    row := r.DB.QueryRow(ctx, 
        "SELECT id, email, password_hash, created_at FROM users WHERE email = $1", email)
    user := &User{}
    if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt); err != nil {
        return nil, err
    }
    return user, nil
}
```
- Use parameterized queries (`$1, $2…`) to prevent SQL injection
- Errors from repos propagate to services unchanged — services don't catch
- Handlers decide HTTP status based on error type or context

## Code Conventions
- **Constructors**: `NewXxx(deps) *Xxx` (no error return, panic if critical validation fails)
- **Context**: All DB/Service calls pass `c.Request.Context()` from handler
- **Error flow**: Repos → Services (wrap for domain logic) → Handlers (translate to HTTP status)
- **Authentication**: JWT HS256 using `JWT_SECRET`, token claims: `{user_id, exp}`
- **Passwords**: Bcrypt with default cost (`bcrypt.DefaultCost ≈ 10`)
- **OTP**: 6 alphanumeric chars, hashed before DB storage, 10-minute expiry, regenerate on each send
- **Email verification**: Check `user != nil` before sending OTP to ensure email not already registered
- **Imports order**: stdlib → third-party (`github.com`) → internal (`clothing-store-backend`)

## Key Files & Patterns
| File | Purpose |
|------|---------|
| `cmd/api/main.go` | Entry: load config, init DB pool, setup router, start server |
| `internal/router/router.go` | Register routes with DI: create Handler/Service/Repo instances, apply middleware |
| `internal/auth/handler.go` | HTTP handlers: SendOTP, ValidateOTP (incomplete), Login |
| `internal/auth/service.go` | Business logic: OTP generation/hashing, email verification, bcrypt login |
| `internal/auth/repository.go` | SQL queries: CreateUser, GetUserByEmail, CreateEmailVerification |
| `internal/auth/jwt.go` | GenerateToken (HS256), RateLimiter (100 req/min by IP) |
| `internal/middleware/jwt.go` | JWTAuth: parse Bearer token, validate signature, extract user_id to context |
| `internal/db/postgres.go` | pgxpool initialization from DATABASE_URL |

## Known Issues & Incomplete Features
- **ValidateOTP handler** (internal/auth/handler.go): Accepts `{email, password}` but has no implementation — complete OTP verification logic before production
- **Rate limiter** (internal/auth/jwt.go): Defined but `SetupRouter()` not called in main.go — currently unused
- **Email sending** (internal/auth/handler.go): Calls `sendEmail()` function — must implement SMTP or email service integration
- **User role field**: Model has `Role` field but schema doesn't define it — either remove or add migration</content>
<parameter name="filePath">d:\Clothing Store\Clothing-Store\Backend\Go\.github\copilot-instructions.md