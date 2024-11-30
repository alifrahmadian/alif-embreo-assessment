package configs

import (
	"database/sql"
	"fmt"

	"github.com/alifrahmadian/alif-embreo-assessment/internal/db"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/handlers"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/repositories"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/services"
)

type Config struct {
	DB       *sql.DB
	Auth     *AuthConfig
	Handlers *Handlers
	Env      string
}

type Handlers struct {
	AuthHandler *handlers.AuthHandler
}

type AuthConfig struct {
	TTL       int
	SecretKey string
}

func LoadConfig() (*Config, error) {
	err := loadGoDotEnv()
	if err != nil {
		return nil, err
	}

	dbConfig := loadDBConfig()
	env := loadEnv()
	authConfig := loadAuthConfig()

	db, err := db.Connect(*dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	// defer db.Close()
	fmt.Println("Database connected successfully: ", db)

	userRepo := repositories.NewUserRepository(db)

	authService := services.NewAuthService(userRepo)

	authHandler := handlers.NewAuthHandler(&authService)

	return &Config{
		DB:   db,
		Env:  env,
		Auth: authConfig,
		Handlers: &Handlers{
			AuthHandler: authHandler,
		},
	}, nil
}
