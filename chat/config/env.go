package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}
	if err := godotenv.Load(fmt.Sprintf(".env.%s", env)); err != nil {
		panic("Error loading .env file")
	}
}