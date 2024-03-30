package config

import (
	"github.com/joho/godotenv"
	"os"
)

func Load() error {
	return godotenv.Load()
}

func Get(key string) string {
	return os.Getenv(key)
}
