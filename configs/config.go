package configs

import (
	"database/sql"
	"fmt"

	"github.com/alifrahmadian/alif-embreo-assessment/internal/db"
)

type Config struct {
	DB       *sql.DB
	Auth     *AuthConfig
	Handlers *Handlers
	Env      string
}

type Handlers struct {
}

type AuthConfig struct {
}

func LoadConfig() (*Config, error) {
	err := loadGoDotEnv()
	if err != nil {
		return nil, err
	}

	dbConfig := loadDBConfig()
	env := loadEnv()

	db, err := db.Connect(*dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	defer db.Close()
	fmt.Println("Database connected successfully")

	return &Config{
		DB:  db,
		Env: env,
	}, nil
}
