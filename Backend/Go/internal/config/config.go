package config

import (
	"os"
)

type Config struct {
	AppPort   string
	DBUrl     string
	JWTSecret string
}

func Load() *Config {
	return &Config{
		AppPort:   getEnv("APP_PORT", "8080"),
		DBUrl:     getEnv("DATABASE_URL", ""),
		JWTSecret: getEnv("JWT_SECRET", "supersecret"),
	}
}

func getEnv(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
