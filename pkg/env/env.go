package env

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnv(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}

func SetupEnvFile() {
	_ = godotenv.Load()
}
