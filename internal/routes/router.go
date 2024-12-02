package routes

import (
	"github.com/alifrahmadian/alif-embreo-assessment/configs"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handlers *configs.Handlers) {
	router.POST("/register", handlers.AuthHandler.Register)
	router.POST("/login", handlers.AuthHandler.Login)
	router.POST("/events/create", handlers.EventHandler.CreateEvent)
	router.GET("/events/:id", handlers.EventHandler.GetEventByID)
	router.GET("/events", handlers.EventHandler.GetAllEvents)
}
