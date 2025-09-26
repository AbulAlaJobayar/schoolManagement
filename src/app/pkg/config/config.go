package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GoEnv    string
	Port     string
	Database string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading environment variables")
	}
	AppConfig = &Config{
		GoEnv:    getEnv("GO_ENV", "development"),
		Port:     getEnv("PORT", "3000"),
		Database: getEnv("DATABASE", ""),
	}
}

func getEnv(key, def string) string {
	value, exist := os.LookupEnv(key)
	if exist {
		return value
	}
	return def
}
