package service

import (
	"github.com/sony/gobreaker"
	"shopdeck.com/messaging_service/app/domain/dto/httprequest"
)

type MessageService interface {
	IsCircuitBreakerEnabled() bool
	GetCircuitBreakerState() gobreaker.State

	GetPrimaryProducerHealth() bool
	GetFallbackProducerHealth() bool

	SendMessage(request httprequest.PublishMessageRequest) error
	SendMessageAsync(request httprequest.PublishMessageRequest)
}
