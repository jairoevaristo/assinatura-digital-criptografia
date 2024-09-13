package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvs() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
