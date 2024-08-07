//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"shopdeck.com/messaging_service/app/controller"
	"shopdeck.com/messaging_service/app/service"
)

// MessageService
var kafkaMessageServiceSet = wire.NewSet(service.KafkaMessageServiceInit)

var defaultMessageServiceSet = wire.NewSet(service.DefaultMessageServiceInit,
	wire.Bind(new(service.MessageService), new(*service.DefaultMessageServiceImpl)))

// Controller
var messageControllerSet = wire.NewSet(controller.MessageControllerInit,
	wire.Bind(new(controller.MessageController), new(*controller.MessageControllerImpl)))

var healthControllerSet = wire.NewSet(controller.HealthControllerInit,
	wire.Bind(new(controller.HealthController), new(*controller.HealthControllerImpl)))

func Init() *Initialization {
	wire.Build(NewInitialization, kafkaMessageServiceSet, defaultMessageServiceSet, messageControllerSet, healthControllerSet)
	return nil
}
