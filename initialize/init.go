package initialize

import (
	"shopdeck.com/messaging_service/app/controller"
	"shopdeck.com/messaging_service/app/service"
	"shopdeck.com/messaging_service/app/service/circuit_breaker"
	"shopdeck.com/messaging_service/app/service/message_producer"
	"shopdeck.com/messaging_service/config"
)

type Initialization struct {
	CircuitBreaker                  circuit_breaker.CircuitBreaker
	PrimaryMessageProducerProvider  message_producer.PrimaryMessageProducerProvider
	FallbackMessageProducerProvider message_producer.FallbackMessageProducerProvider
	MessageService                  service.MessageService
	MessageController               controller.MessageController
	HealthController                controller.HealthController
}

func Init() *Initialization {
	configuration := config.GetConfiguration()

	// MessageService
	circuitBreaker := circuit_breaker.CircuitBreakerInit(configuration)
	var primaryMessageProducerProvider = message_producer.PrimaryMessageProducerProviderInit(configuration)
	var fallbackMessageProducerProvider = message_producer.FallbackMessageProducerProviderInit(configuration)
	var messageService = service.MessageServiceImplInit(primaryMessageProducerProvider, fallbackMessageProducerProvider, circuitBreaker)

	// Controller
	var messageControllerSet = controller.MessageControllerInit(messageService)
	var healthControllerSet = controller.HealthControllerInit(messageService)

	return &Initialization{
		CircuitBreaker:                  circuitBreaker,
		PrimaryMessageProducerProvider:  primaryMessageProducerProvider,
		FallbackMessageProducerProvider: fallbackMessageProducerProvider,
		MessageService:                  messageService,
		MessageController:               messageControllerSet,
		HealthController:                healthControllerSet,
	}
}
