package config

import (
	"shopdeck.com/messaging_service/app/controller"
	"shopdeck.com/messaging_service/app/service"
)

type Initialization struct {
	MessageService    service.MessageService
	MessageController controller.MessageController
}

func NewInitialization(messageService service.MessageService,
	messageController controller.MessageController) *Initialization {
	return &Initialization{
		MessageService:    messageService,
		MessageController: messageController,
	}
}
