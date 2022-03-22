package secrets

import (
	"os"
)

func FromEnv(key string, fallback string) string {
	value, found := os.LookupEnv("DATABASE_URL")
	if found {
		return value
	}
	return fallback
}
