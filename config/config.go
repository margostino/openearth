package config

import (
	"github.com/joho/godotenv"
	"os"
)

var env = loadEnv()

func IsDevEnv() bool {
	return env == "development"
}

func loadEnv() string {
	godotenv.Load()
	return os.Getenv("VERCEL_ENV")
}
