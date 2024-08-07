package router

import (
	"github.com/gin-gonic/gin"
	"shopdeck.com/messaging_service/config"
)

func Init(init *config.Initialization) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		message := api.Group("/message")
		message.POST("", init.MessageController.PublishMessage)
		message.GET("/start", init.MessageController.StartTest)
	}

	return router
}
