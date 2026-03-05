package config

import (
	"log"
	"os"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	Port       string
}

func getenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environmental variable %s is not set", key)
	}
	return value
}

func LoadConfig() *Config{
	return &Config{
		DBHost: getenv("DB_HOST"),
		DBUser: getenv("DB_USER"),
		DBPassword: getenv("DB_PASSWORD"),
		DBName: getenv("DB_NAME"),
		DBPort: getenv("DB_PORT"),
		Port: getenv("PORT"),
	}
}