package configs

import (
	"fmt"
	"os"

	"github.com/alifrahmadian/alif-embreo-assessment/internal/db"
	"github.com/joho/godotenv"
)

func loadDBConfig() *db.DBConfig {
	return &db.DBConfig{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func loadEnv() string {
	return os.Getenv("ENV")
}

func loadGoDotEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	return nil
}
