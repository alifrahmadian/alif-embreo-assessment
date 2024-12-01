package main

import (
	"fmt"

	"github.com/alifrahmadian/alif-embreo-assessment/configs"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/db"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/handlers"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/repositories"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/routes"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/services"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Config *configs.Config
}

func LoadConfig() (*configs.Config, error) {
	err := configs.LoadGoDotEnv()
	if err != nil {
		return nil, err
	}

	dbConfig := configs.LoadDBConfig()
	env := configs.LoadEnv()
	authConfig := configs.LoadAuthConfig()

	db, err := db.Connect(*dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	// defer db.Close()
	fmt.Println("Database connected successfully: ", db)
	fmt.Println(authConfig)

	userRepo := repositories.NewUserRepository(db)

	authService := services.NewAuthService(userRepo)

	authHandler := handlers.NewAuthHandler(&authService, authConfig.SecretKey, authConfig.TTL)

	return &configs.Config{
		DB:   db,
		Env:  env,
		Auth: authConfig,
		Handlers: &configs.Handlers{
			AuthHandler: authHandler,
		},
	}, nil
}

func NewApp() *App {
	cfg, err := LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	router := gin.Default()
	routes.SetupRoutes(router, cfg.Handlers)

	return &App{
		Router: router,
		Config: cfg,
	}
}
