package config

import (
	"os"
)

var (
	DatabaseName     = getEnv("DB_DATABASE", "go-short")
	DatabaseUsername = getEnv("DB_USERNAME", "root")
	DatabasePasword  = getEnv("DB_PASSWORD", "poot")
	DatabaseHost     = getEnv("DB_HOST", "127.0.0.1")
	DatabasePort     = getEnv("DB_PORT", "3306")
	Port             = getEnv("PORT", "8080")
)

func getEnv(key string, fallback string) string {
	env := os.Getenv(key)

	if env == "" {
		return fallback
	}
	return env
}
