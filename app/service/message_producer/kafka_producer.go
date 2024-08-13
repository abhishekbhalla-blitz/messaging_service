package message_producer

import (
	"errors"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
	"shopdeck.com/messaging_service/app/domain/dto/httprequest"
	"shopdeck.com/messaging_service/config"
	"time"
)

type KafkaMessageProducer struct {
	healthTopic string
	producer    *kafka.Producer
}

func KafkaMessageProducerInit(configuration config.KafkaConfig) MessageProducer {
	producerConfig := getKafkaConfigMap(configuration.GetBrokers())

	producer, producerError := kafka.NewProducer(&producerConfig)
	if producerError != nil {
		log.Error("Unable to create producer. Error: ", producerError)
	}

	return &KafkaMessageProducer{
		healthTopic: configuration.HealthCheckTopic,
		producer:    producer,
	}
}

func getKafkaConfigMap(servers string) kafka.ConfigMap {
	return kafka.ConfigMap{
		"bootstrap.servers":       servers, // OVERRIDDEN, csv brokers list"
		"acks":                    "all",   // POSSIBLY OVERRIDDEN, all ISRs should acknowledge the message
		"retries":                 5,       // default
		"linger.ms":               100,     // increased from 10 to improve throughput and reduce load
		"batch.size":              16384,   // default
		"batch.num.messages":      50,      // set by trial and error
		"compression.type":        "gzip",  // compression to improve throughput, increase cpu load on server
		"message.timeout.ms":      30000,   // 30 seconds
		"enable.idempotence":      true,    // ensure exactly-once delivery
		"go.delivery.reports":     true,    // enable delivery reports
		"socket.keepalive.enable": true,    // keep connection alive
	}
}

func (kmp *KafkaMessageProducer) SendMessage(request httprequest.PublishMessageRequest) error {
	message := extractMessage(request)
	deliveryChan := make(chan kafka.Event, 1)
	return kmp.sendMessage(message, deliveryChan)
}

func extractMessage(request httprequest.PublishMessageRequest) kafka.Message {
	return kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &request.Target,
			Partition: kafka.PartitionAny,
		},
		Key:           []byte(""),
		Value:         []byte(request.Message),
		TimestampType: kafka.TimestampLogAppendTime,
	}
}

func (kmp *KafkaMessageProducer) sendMessage(message kafka.Message, deliveryChan chan kafka.Event) error {
	err := kmp.producer.Produce(&message, deliveryChan)
	if err != nil {
		return err
	}

	select {
	case e := <-deliveryChan:
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			log.Printf("Delivery failed: %v", m.TopicPartition.Error)
			return m.TopicPartition.Error
		}
	case <-time.After(3000*time.Millisecond + 100*time.Millisecond):
		log.Println("Timed out waiting for delivery report")
		return errors.New("timed out waiting for delivery report")
	}

	return nil
}

func (kmp *KafkaMessageProducer) HealthCheck() bool {
	if kmp.producer == nil {
		return false
	}

	var healthy = false
	metadata, err := kmp.producer.GetMetadata(&kmp.healthTopic, false, 1000)
	if err != nil {
		log.Error("Unable to verify producer. Error: ", err)
	} else if metadata == nil {
		log.Error("Unable to verify producer. Error: ", err)
	} else if len(metadata.Brokers) == 0 {
		log.Error("No brokers available for health check topic")
	} else {
		healthy = true
	}

	return healthy
}
