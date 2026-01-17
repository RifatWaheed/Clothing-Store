package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	// "golang.org/x/text/message"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(ctx context.Context, email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo.CreateUser(ctx, email, string(hash))
}

func (s *Service) Login(ctx context.Context, email, password string) (*User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *Service) VerifyEmailExistence(ctx context.Context, email string) *User {
	user, _ := s.repo.GetUserByEmail(ctx, email)
	return user
}

func (s *Service) SendOTPToEmail(ctx context.Context, otp string, email string) (string, error) {
	message := ""
	// err := s.repo.SendOTPToEmail(ctx, otp, email)
	return message, nil
}
