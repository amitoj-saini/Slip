package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file (might not exist)")
	}
}

func GetEnv(env string, defaultString string) string {
	env_var := os.Getenv(env)
	if env_var == "" {
		return defaultString
	}

	return env_var
}