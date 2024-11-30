package routes

import (
	"github.com/alifrahmadian/alif-embreo-assessment/configs"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handlers *configs.Handlers) {
	router.POST("/register", handlers.AuthHandler.Register)
}
