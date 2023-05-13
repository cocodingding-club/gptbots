package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadConfig() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	err := godotenv.Load(".env" + env)
	if err != nil {
		log.Fatalf("Failed to load .env.%s file", env)
	}
}
