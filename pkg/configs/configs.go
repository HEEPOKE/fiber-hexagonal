package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Cfg *Config
)

type Config struct {
	PORT       string
	DBHost     string
	DBUserName string
	DBPassword string
	DBName     string
	DBPort     string
	DBssl      string
	DBTimezone string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	config := &Config{
		PORT:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBUserName: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBssl:      os.Getenv("DB_SSL"),
		DBTimezone: os.Getenv("DB_TIMEZONE"),
	}

	Cfg = config

	return config, nil
}
