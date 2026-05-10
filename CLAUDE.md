# Clothing Store — Claude Code Guide

## Project Overview

An online clothing store built with **Angular 21.1** (frontend), **Go / Gin** (backend), and **PostgreSQL** (database). The project is in active development.

---

## Tech Stack

| Layer     | Technology                                                   |
|-----------|--------------------------------------------------------------|
| Frontend  | Angular 21.1, Angular Material 21.1, TypeScript 5.9, SSR    |
| Backend   | Go 1.25, Gin v1.11, pgx/v5, golang-jwt/v5                   |
| Database  | PostgreSQL, golang-migrate                                    |
| Email     | SendGrid (prod) / Mock (dev)                                 |
| Auth      | OTP via email → JWT (24 h expiry)                            |
| Testing   | Vitest (frontend), Go test (backend)                         |

---

## Repository Structure

```
Clothing-Store/
├── Frontend/Angular/
│   └── src/app/
│       ├── models/          # TypeScript data models
│       ├── services/        # Angular injectable services
│       └── views/           # All UI components
│           ├── navbar/
│           ├── landing/
│           ├── about-us/
│           ├── public-layout/
│           ├── global-footer/
│           ├── auth/
│           │   └── register/
│           └── admin/
│               ├── admin-shell/
│               ├── admin-dashboard/
│               ├── products/
│               ├── orders/
│               └── inventory/
└── Backend/Go/
    ├── cmd/
    │   ├── api/main.go      # Server entry point
    │   └── migrate/main.go  # Migration runner
    ├── internal/
    │   ├── auth/            # Handler, Service, Repository, Model, JWT, Utility
    │   ├── product/         # Handler, Service, Repository, Model
    │   ├── middleware/       # JWT auth middleware
    │   ├── email/           # SendGrid + Mock implementations
    │   ├── config/          # Env config loader
    │   ├── db/              # PostgreSQL pool init
    │   └── router/          # Gin route setup
    └── migrations/          # SQL up/down migration files
```

---

## Backend API Routes

### Public (no auth)
| Method | Path | Description |
|--------|------|-------------|
| GET | `/health` | Health check |
| POST | `/api/auth/request-sendOTP` | Send OTP to email |
| POST | `/api/auth/request-validateOTP` | Validate OTP |
| POST | `/api/auth/login` | Login, returns JWT |
| GET | `/api/public/products` | List products (paginated) |
| GET | `/api/public/products/:id` | Get product by ID |

### Protected (JWT required)
| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/me` | Returns authenticated user_id |
| GET | `/api/products` | List products (auth) |
| GET | `/api/products/:id` | Get product by ID (auth) |

---

## Frontend Routes

| Path | Component | Notes |
|------|-----------|-------|
| `/` | Landing | Public layout |
| `/about-us` | AboutUs | Public layout |
| `/register` | Register | Public layout |
| `/admin` | AdminShell | Redirects to dashboard |
| `/admin/admin-dashboard` | AdminDashboard | |
| `/admin/orders` | Orders | |
| `/admin/products` | Products | |
| `/admin/inventory` | Inventory | |

---

## Database Schema (current)

**users** — UUID PK, email (unique), password_hash, created_at  
**products** — BIGSERIAL PK, name, price, discount, SKU, color, gender, size, stock, type, voucher  
**email_verifications** — OTP storage for auth flow  

---

## Development Status

### Done
- Full backend API: auth (OTP + JWT), products CRUD
- Frontend routing and layout structure
- Navbar, Register form, Admin shell skeleton
- DB schema + migrations
- Email service (SendGrid + Mock)

### In Progress / TODO
- `ProductService` (Angular) — currently empty, needs API integration
- Admin pages: Products, Orders, Inventory — empty shells
- Login page — missing from frontend
- Auth guard for admin routes
- Frontend ↔ Backend API wiring

---

## Working Rules (important)

1. **Always ask before editing any file.** Show the proposed change and wait for approval.
2. **Only write code when at least 95% confident it will work.** If unsure, say so and explain the uncertainty before writing anything.
3. **Maintain folder structure and code hygiene.** New files go in the correct directory; naming must match existing conventions (kebab-case folders, PascalCase Angular class names, snake_case Go identifiers).
4. **No unsolicited refactors, abstractions, or feature additions.** Implement exactly what is asked, nothing more.
5. **Suggest improvements proactively, but do not implement them without approval.** If there is a clearly better approach, flag it as a suggestion before proceeding.
6. **No comments unless the WHY is non-obvious.** Do not add explanatory comments for self-evident code.

---

## Angular Conventions

- Standalone components only (no NgModules)
- Signal-based reactivity (`signal()`, `computed()`)
- Angular Material for all UI components
- Services go in `src/app/services/<domain>/<domain>.service.ts`
- Components go in `src/app/views/<feature>/<feature>.ts` + `<feature>.html` + `<feature>.scss`
- Models go in `src/app/models/<domain>/<domain>.ts`

## Go Conventions

- Layered architecture per domain: `handler.go → service.go → repository.go → model.go`
- No global state; dependencies injected via constructors
- All DB access through pgx pool (no ORM)
- New domain packages go under `internal/<domain>/`
- Migrations go in `migrations/` as numbered SQL files
