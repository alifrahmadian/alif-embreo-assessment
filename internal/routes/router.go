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
	router.PUT("/events/:id/approve", handlers.EventHandler.ApproveEvent)
	router.PUT("/events/:id/reject", handlers.EventHandler.RejectEvent)
	router.GET("/events", handlers.EventHandler.GetAllEvents)
	router.GET("/vendors", handlers.VendorHandler.GetVendors)
}
