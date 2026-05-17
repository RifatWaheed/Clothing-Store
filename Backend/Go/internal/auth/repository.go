package auth

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool" //A pool of reusable database connections for postgres
)

type Repository struct {
	DB *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository { // works similar as a constructor
	return &Repository{DB: db}
}

func (r *Repository) CreateUser(ctx context.Context, email, passwordHash, name string) error {
	_, err := r.DB.Exec(ctx,
		"INSERT INTO users (email, password_hash, name) VALUES ($1, $2, $3)",
		email, passwordHash, name)
	return err
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	row := r.DB.QueryRow(
		ctx,
		`SELECT id, email, password_hash, created_at, role FROM users WHERE email = $1`,
		email,
	)

	user := User{}
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) CreateEmailVerification(ctx context.Context, email string, otpHash string, expiresAt time.Time) error {
	_, err := r.DB.Exec(ctx, `
		INSERT INTO email_verifications (email, otp_hash, expires_at)
		VALUES ($1, $2, $3)
	`, email, otpHash, expiresAt)
	return err
}

func (r *Repository) GetActiveOTP(ctx context.Context, email string) (*EmailVerification, error) {
	row := r.DB.QueryRow(ctx,
		`SELECT id, email, otp_hash, expires_at FROM email_verifications
		 WHERE email = $1 AND expires_at > NOW()`,
		email,
	)
	var ev EmailVerification
	if err := row.Scan(&ev.ID, &ev.Email, &ev.OTPHash, &ev.ExpiresAt); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (r *Repository) DeleteOTPByEmail(ctx context.Context, email string) error {
	_, err := r.DB.Exec(ctx, `DELETE FROM email_verifications WHERE email = $1`, email)
	return err
}

func (r *Repository) DeleteExpiredOTPs(ctx context.Context) error {
	_, err := r.DB.Exec(ctx, `DELETE FROM email_verifications WHERE expires_at < NOW()`)
	return err
}

func (r *Repository) SendOTPToEmail(ctx context.Context, otp string, email string) {
	// saveOTPandEmailInfoToDb()
}
