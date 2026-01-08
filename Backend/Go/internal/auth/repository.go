package auth

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool" //A pool of reusable database connections for postgres
)

type Repository struct {
	DB *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository { // works similar as a constructor
	return &Repository{DB: db}
}

func (r *Repository) CreateUser(ctx context.Context, email, passwordHash string) error {
	_, err := r.DB.Exec(ctx, "INSERT into users (email, password_hash) VALUES ($1 , $2)", email, passwordHash)
	return err
}

func (r *Repository) GetUserByEmail(ctx context.Context, email, passwordHash string) (*User, error) {
	row := r.DB.QueryRow(ctx, `SELECT id, email, password_hash, created_at FROM users WHERE email = $1`, email)
	user := User{}
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
