package config

import (
	"shopdeck.com/messaging_service/app/controller"
	"shopdeck.com/messaging_service/app/service"
)

type Initialization struct {
	MessageService    service.MessageService
	MessageController controller.MessageController
	HealthController  controller.HealthController
}

func NewInitialization(messageService service.MessageService,
	messageController controller.MessageController,
	healthController controller.HealthController) *Initialization {
	return &Initialization{
		MessageService:    messageService,
		MessageController: messageController,
		HealthController:  healthController,
	}
}
