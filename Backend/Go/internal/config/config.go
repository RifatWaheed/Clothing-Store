package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// type Config struct {
// 	AppPort   string
// 	DBUrl     string
// 	JWTSecret string
// }

type Config struct {
	DatabaseURL string
}

func Load() *Config {
	// Load .env file (optional, will ignore if not found)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
