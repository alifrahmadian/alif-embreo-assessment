package main

import (
	config "github.com/alifrahmadian/alif-embreo-assessment/configs"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/routes"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Config *config.Config
}

func NewApp() *App {
	cfg, err := config.LoadConfig()
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
