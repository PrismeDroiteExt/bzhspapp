package config

import (
	"os"
	"time"
)

type Config struct {
	JWTSecret          string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
	PasswordHashCost   int
}

func LoadConfig() *Config {
	return &Config{
		JWTSecret:          os.Getenv("JWT_SECRET"),
		AccessTokenExpiry:  24 * time.Hour,
		RefreshTokenExpiry: 7 * 24 * time.Hour,
		PasswordHashCost:   12,
	}
}
