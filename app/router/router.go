package router

import (
	"github.com/gin-gonic/gin"
	"shopdeck.com/messaging_service/config"
	"shopdeck.com/messaging_service/initialize"
)

func Init(init *initialize.Initialization) *gin.Engine {
	configuration := config.GetConfiguration()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	messagingService := router.Group(configuration.Server.ContextPath)

	// health route group
	// can add more details for more comprehensive health check
	messagingService.GET("/health", init.HealthController.ServiceHealthCheck)

	// api route group
	api := messagingService.Group("/api")
	{
		// message route group
		message := api.Group("/message")
		message.POST("", init.MessageController.SendMessage)
		message.GET("/start", init.MessageController.StartTest)
	}

	return router
}
