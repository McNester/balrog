package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("No discovery file!")
	}
}

func GetEnvVariable(key string) string {
	return os.Getenv(key)
}
