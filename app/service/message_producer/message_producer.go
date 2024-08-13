package message_producer

import "shopdeck.com/messaging_service/app/domain/dto/httprequest"

type MessageProducer interface {
	SendMessage(request httprequest.PublishMessageRequest) error
	HealthCheck() bool
}
