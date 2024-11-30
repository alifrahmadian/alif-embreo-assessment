package main

import (
	config "github.com/alifrahmadian/alif-embreo-assessment/configs"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Config *config.Config
}
