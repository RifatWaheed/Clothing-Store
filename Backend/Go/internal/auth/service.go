package auth

import (
	emailpkg "clothing-store-backend/internal/email"
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo        *Repository
	emailSender emailpkg.EmailSender
}

func NewService(repo *Repository, emailSender emailpkg.EmailSender) *Service {
	return &Service{
		repo:        repo,
		emailSender: emailSender,
	}
}

func (s *Service) Register(ctx context.Context, email, password, name string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo.CreateUser(ctx, email, string(hash), name)
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
	existing, err := s.repo.GetActiveOTP(ctx, email)
	if err == nil && existing != nil {
		return errors.New("OTP already sent to this email, please check your inbox")
	}

	// clear stale expired record if present before inserting new one
	_ = s.repo.DeleteOTPByEmail(ctx, email)

	otp, err := generateAlphaNumericOTP(6)
	if err != nil {
		return err
	}

	otpHash, err := hashOTP(otp)
	if err != nil {
		return err
	}

	expiresAt := time.Now().Add(10 * time.Minute)

	if err := s.repo.CreateEmailVerification(ctx, email, otpHash, expiresAt); err != nil {
		return err
	}

	html, err := emailpkg.RenderOTPEmail(otp, email)
	if err != nil {
		return err
	}

	return s.emailSender.Send(ctx, email, "Your verification code", html)
}

func (s *Service) ValidateOTP(ctx context.Context, email, otp string) error {
	record, err := s.repo.GetActiveOTP(ctx, email)
	if err != nil {
		return errors.New("no valid OTP found for this email")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(record.OTPHash), []byte(otp)); err != nil {
		return errors.New("invalid OTP")
	}

	return s.repo.DeleteOTPByEmail(ctx, email)
}
