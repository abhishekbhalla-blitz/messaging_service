package main

import (
	"shopdeck.com/messaging_service/app/router"
	"shopdeck.com/messaging_service/config"
	"shopdeck.com/messaging_service/initialize"
	"strconv"
)

func init() {
	config.InitConfig()
	initialize.InitLog()
}

func main() {
	configuration := config.GetConfiguration()
	init := initialize.Init()
	app := router.Init(init)
	app.Run(":" + strconv.Itoa(configuration.Server.Port))
}
