package service

import (
	httprequest "shopdeck.com/messaging_service/app/domain/dto"
)

type MessageService interface {
	PublishMessage(request httprequest.PublishMessageRequest) error
	PublishMessageAsync(request httprequest.PublishMessageRequest) error
}

func MessageServiceInit(messageService DefaultMessageServiceImpl) *DefaultMessageServiceImpl {
	return &messageService
}
