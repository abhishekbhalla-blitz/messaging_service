package service

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/sony/gobreaker"
	"shopdeck.com/messaging_service/app/domain/dto/httprequest"
	"shopdeck.com/messaging_service/app/service/circuit_breaker"
	"shopdeck.com/messaging_service/app/service/message_producer"
)

type MessageServiceImpl struct {
	primaryProducer  message_producer.MessageProducer
	fallbackProducer message_producer.MessageProducer
	circuitBreaker   circuit_breaker.CircuitBreaker
}

func MessageServiceImplInit(primaryProvider message_producer.PrimaryMessageProducerProvider,
	fallbackProvider message_producer.FallbackMessageProducerProvider,
	circuitBreaker circuit_breaker.CircuitBreaker) *MessageServiceImpl {

	if primaryProvider.MessageProducer == nil {
		log.Fatal("PrimaryProducer is required")
	}

	return &MessageServiceImpl{
		primaryProducer:  primaryProvider.MessageProducer,
		fallbackProducer: fallbackProvider.MessageProducer,
		circuitBreaker:   circuitBreaker,
	}
}

func (messageService *MessageServiceImpl) IsCircuitBreakerEnabled() bool {
	return messageService.circuitBreaker.Breaker != nil
}

func (messageService *MessageServiceImpl) GetCircuitBreakerState() gobreaker.State {
	var cbState gobreaker.State
	if messageService.IsCircuitBreakerEnabled() {
		cbState = messageService.circuitBreaker.Breaker.State()
	} else {
		cbState = gobreaker.StateClosed
	}
	return cbState
}

func (messageService *MessageServiceImpl) SendMessageAsync(request httprequest.PublishMessageRequest) {
	go func() {
		err := messageService.SendMessage(request)
		if err != nil {
			log.Printf("Error sending message asynchronously: %v", err)
		}
	}()
}

func (messageService *MessageServiceImpl) SendMessage(request httprequest.PublishMessageRequest) error {
	if !messageService.IsCircuitBreakerEnabled() {
		return messageService._sendMessage(messageService.primaryProducer, request)
	} else {
		_, pErr := messageService.circuitBreaker.Breaker.Execute(func() (interface{}, error) {
			// Attempt to send message to primary cluster
			if err := messageService._sendMessage(messageService.primaryProducer, request); err != nil {
				return nil, err
			}
			return nil, nil
		})

		if pErr != nil {
			log.Error("Failed to send to primary cluster: %v\nAttempting to send message to fallback cluster.", pErr)
			if fErr := messageService._sendMessage(messageService.fallbackProducer, request); fErr != nil {
				log.Error("Failed to send to fallback cluster: %v\n", fErr)
			}
			return errors.New("unable to send message")
		}
	}

	return nil
}

func (messageService *MessageServiceImpl) _sendMessage(producer message_producer.MessageProducer,
	message httprequest.PublishMessageRequest) error {
	return producer.SendMessage(message)
}

func (messageService *MessageServiceImpl) GetPrimaryProducerHealth() bool {
	return messageService.primaryProducer.HealthCheck()
}

func (messageService *MessageServiceImpl) GetFallbackProducerHealth() bool {
	return messageService.fallbackProducer.HealthCheck()
}
