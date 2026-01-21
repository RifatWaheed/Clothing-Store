package config

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env file at package initialization
	_ = godotenv.Load()
}

// type Config struct {
// 	AppPort   string
// 	DBUrl     string
// 	JWTSecret string
// }

type Config struct {
	DatabaseURL    string
	JWTSecret      string
	ServerPort     string
	SendGridAPIKey string
	SendGridFrom   string
}

func Load() *Config {
	return &Config{
		DatabaseURL:    os.Getenv("DATABASE_URL"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		ServerPort:     os.Getenv("SERVER_PORT"),
		SendGridAPIKey: os.Getenv("SENDGRID_API_KEY"),
		SendGridFrom:   os.Getenv("SENDGRID_FROM_EMAIL"),
	}
}
