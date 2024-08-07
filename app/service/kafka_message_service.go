package service

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
	httprequest "shopdeck.com/messaging_service/app/domain/dto"
)

type KafkaMessageService struct {
	producer *kafka.Producer
}

func KafkaMessageServiceInit() *KafkaMessageService {
	var p, _ = kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers":  "my-cluster-kafka-bootstrap:9092",
			"linger.ms":          100,
			"batch.num.messages": 50,
			//"enable.idempotence": true,
			//"retries":            3,
			//"acks":               "all",
			"compression.type": "gzip",
			//"request.timeout.ms": 30000,
		})
	return &KafkaMessageService{
		// Initialize Kafka producer and consumer.
		//producer := kafka.NewSyncProducer(kafkaBrokers)
		// consumer := kafka.NewConsumer(kafkaBrokers, kafkaConsumerConfig)
		producer: p,
	}
}

func (messageService KafkaMessageService) PublishMessage(request httprequest.PublishMessageRequest) error {
	//log.Info(request.Message)
	message := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &request.Target, Partition: kafka.PartitionAny},
		Value:          []byte(request.Message),
		TimestampType:  kafka.TimestampLogAppendTime,
	}

	err := messageService.producer.Produce(&message, nil)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (messageService KafkaMessageService) PublishMessageAsync(request httprequest.PublishMessageRequest) error {
	log.Info(request.Message)
	return nil
}
