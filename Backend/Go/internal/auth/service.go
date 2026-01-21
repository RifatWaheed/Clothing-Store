package auth

import (
	"clothing-store-backend/internal/email"
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo        *Repository
	emailSender email.EmailSender
}

func NewService(repo *Repository, emailSender email.EmailSender) *Service {
	return &Service{
		repo:        repo,
		emailSender: emailSender,
	}
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

func (s *Service) SendOTPToEmail(ctx context.Context, email string) error {

	otp, errorGenerateOTP := generateAlphaNumericOTP(6) // OTP length is currently 6
	if errorGenerateOTP != nil {
		return errorGenerateOTP
	}

	otpHash, err := hashOTP(otp)
	if err != nil {
		return err
	}

	expiresAt := time.Now().Add(10 * time.Minute)

	if err := s.repo.CreateEmailVerification(ctx, email, otpHash, expiresAt); err != nil {
		return err
	}

	return s.emailSender.Send(
		ctx,
		email,
		"Your verification code",
		"Your OTP is: "+otp,
	)
}
