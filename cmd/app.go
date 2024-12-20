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
	eventRepo := repositories.NewEventRepository(db)
	vendorRepo := repositories.NewVendorRepository(db)

	authService := services.NewAuthService(userRepo)
	eventService := services.NewEventService(eventRepo)
	vendorService := services.NewVendorService(vendorRepo)

	authHandler := handlers.NewAuthHandler(&authService, authConfig.SecretKey, authConfig.TTL)
	eventHandler := handlers.NewEventHandler(&eventService)
	vendorHandler := handlers.NewVendorHandler(&vendorService)

	return &configs.Config{
		DB:   db,
		Env:  env,
		Auth: authConfig,
		Handlers: &configs.Handlers{
			AuthHandler:   authHandler,
			EventHandler:  eventHandler,
			VendorHandler: vendorHandler,
		},
	}, nil
}

func NewApp() *App {
	cfg, err := LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	router := gin.Default()
	routes.SetupRoutes(cfg.Auth.SecretKey, router, cfg.Handlers)

	return &App{
		Router: router,
		Config: cfg,
	}
}
