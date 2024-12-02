package routes

import (
	"github.com/alifrahmadian/alif-embreo-assessment/configs"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/constants"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(secretKey string, router *gin.Engine, handlers *configs.Handlers) {
	router.POST("/register", handlers.AuthHandler.Register)
	router.POST("/login", handlers.AuthHandler.Login)
	router.POST("/events/create", middlewares.AuthMiddleware(secretKey, constants.RoleHR), handlers.EventHandler.CreateEvent)
	router.GET("/events/:id", middlewares.AuthMiddleware(secretKey, constants.RoleHR, constants.RoleVendor), handlers.EventHandler.GetEventByID)
	router.PUT("/events/:id/approve", middlewares.AuthMiddleware(secretKey, constants.RoleVendor), handlers.EventHandler.ApproveEvent)
	router.PUT("/events/:id/reject", middlewares.AuthMiddleware(secretKey, constants.RoleVendor), handlers.EventHandler.RejectEvent)
	router.GET("/events", middlewares.AuthMiddleware(secretKey, constants.RoleHR, constants.RoleVendor), handlers.EventHandler.GetAllEvents)
	router.GET("/vendors", handlers.VendorHandler.GetVendors)
}
