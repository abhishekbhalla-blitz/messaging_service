package message_producer

import (
	log "github.com/sirupsen/logrus"
	"shopdeck.com/messaging_service/config"
)

type PrimaryMessageProducerProvider struct {
	MessageProducer MessageProducer
}

func PrimaryMessageProducerProviderInit(configuration config.Config) PrimaryMessageProducerProvider {
	producer := KafkaMessageProducerInit(configuration.Producer.Primary.ProducerConfig.KafkaConfig)

	healthy := producer.HealthCheck()
	if !healthy {
		log.Error("Primary producer unhealthy")
		if configuration.Producer.Primary.FastFail {
			log.Fatal("Fast fail enabled for primary producer. Quitting.")
		}
	}

	return PrimaryMessageProducerProvider{
		MessageProducer: producer,
	}
}

type FallbackMessageProducerProvider struct {
	MessageProducer MessageProducer
}

func FallbackMessageProducerProviderInit(configuration config.Config) FallbackMessageProducerProvider {
	provider := FallbackMessageProducerProvider{}

	if configuration.Producer.Fallback.Enabled {
		producer := KafkaMessageProducerInit(configuration.Producer.Fallback.ProducerConfig.KafkaConfig)
		provider.MessageProducer = producer

		healthy := producer.HealthCheck()
		if !healthy {
			log.Error("Fallback producer unhealthy")
			if configuration.Producer.Fallback.FastFail {
				log.Fatal("Fast fail enabled for fallback producer. Quitting.")
			}
		}
	}

	return provider
}
