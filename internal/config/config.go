package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env string
	Port string
	DbUrl string
	RedisUrl string
	PasetoKey string
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	
	return  fallback
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Env: getEnv("ENV", "dev"),
		Port: getEnv("PORT", "localhost:8080"),
		DbUrl: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/cinema?sslmode=disable"),
		RedisUrl: getEnv("REDIS_URL", "redis://localhost:6379"),
		PasetoKey: getEnv("PASETO_KEY", ""),
	}

	if cfg.DbUrl == "" {
		return nil, errors.New("DB_URL is required")
	}

	if cfg.RedisUrl == "" {
		return  nil, errors.New("REDIS_URL is required")
	}

	if cfg.PasetoKey == "" {
		return  nil, errors.New("PASETO_KEY is required")
	}

	return  cfg, nil
}