package auth

import "time"

type User struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	Role         string // "customer", "admin"
}

type EmailVerification struct {
	ID        string
	Email     string
	OTPHash   string
	ExpiresAt time.Time
}
