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

// Message Controller
var messageControllerSet = wire.NewSet(controller.MessageControllerInit,
	wire.Bind(new(controller.MessageController), new(*controller.MessageControllerImpl)))

func Init() *Initialization {
	wire.Build(NewInitialization, kafkaMessageServiceSet, defaultMessageServiceSet, messageControllerSet)
	return nil
}
