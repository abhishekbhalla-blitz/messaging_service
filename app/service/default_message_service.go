package service

import (
	"errors"
	"shopdeck.com/messaging_service/app/constant/messageType"
	httprequest "shopdeck.com/messaging_service/app/domain/dto"
)

type DefaultMessageServiceImpl struct {
	kafkaMessageService *KafkaMessageService
}

func DefaultMessageServiceInit(kafkaMessageService *KafkaMessageService) *DefaultMessageServiceImpl {
	return &DefaultMessageServiceImpl{kafkaMessageService}
}

func (messageService DefaultMessageServiceImpl) getMessageHandlerService(mType messageType.MessageType) MessageService {
	var messageHandlerService MessageService = nil

	switch mType {
	case messageType.KAFKA:
		messageHandlerService = messageService.kafkaMessageService
	}

	return messageHandlerService
}

func (messageService DefaultMessageServiceImpl) PublishMessage(request httprequest.PublishMessageRequest) error {
	messageHandlerService := messageService.getMessageHandlerService(request.MessageType)
	if messageHandlerService == nil {
		return errors.New("illegal message type")
		// Implement other message types as needed.
		// return ErrUnsupportedMessageType
	}

	if err := messageHandlerService.PublishMessage(request); err != nil {
		return err
	}

	return nil
}

func (messageService DefaultMessageServiceImpl) PublishMessageAsync(request httprequest.PublishMessageRequest) error {
	messageHandlerService := messageService.getMessageHandlerService(request.MessageType)
	if messageHandlerService != nil {
		if err := messageHandlerService.PublishMessageAsync(request); err != nil {
			return err
		}
	}
	return nil
}
