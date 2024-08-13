package service

//
//type KafkaMessageService struct {
//	circuitBreakerEnabled bool
//	circuitBreaker        *gobreaker.CircuitBreaker
//	admin                 *kafka.AdminClient
//	primaryProducer       *kafka.Producer
//	fallbackProducer      *kafka.Producer
//}

//func KafkaMessageServiceInit() *KafkaMessageService {
//	var kafkaMessageService KafkaMessageService
//	configuration := config.GetConfiguration()
//
//	// admin client
//	admin, adminErr := getAdminClient(configuration.Producer.Primary.ProducerConfig.KafkaConfig.GetBrokers())
//	if adminErr != nil {
//		log.Error("Unable to get connect to primary cluster via admin. Error: ", adminErr)
//	}
//	kafkaMessageService.admin = admin
//
//	// primaryProducer
//	var primaryProducer, primaryProducerError = getProducer(configuration.Producer.Primary.ProducerConfig.KafkaConfig.GetBrokers())
//	if primaryProducer == nil {
//		log.Fatal("Unable to get connect to primary cluster via producer. Error: ", primaryProducerError)
//	}
//	kafkaMessageService.primaryProducer = primaryProducer
//
//	// fallbackProducer, if enabled else nil
//	if configuration.Producer.Fallback.Enabled {
//		kafkaMessageService.circuitBreaker = getCircuitBreaker()
//
//		fallbackProducer, fallbackProducerError := getProducer(configuration.Producer.Fallback.ProducerConfig.KafkaConfig.GetBrokers())
//		if fallbackProducerError != nil {
//			log.Fatal("Unable to get connect to fallback cluster via producer. Error: ", fallbackProducerError)
//		} else {
//			kafkaMessageService.fallbackProducer = fallbackProducer
//		}
//	}
//
//	return &kafkaMessageService
//}
//
////func getCircuitBreaker() *gobreaker.CircuitBreaker {
////	cbSettings := gobreaker.Settings{
////		Name:        "PrimaryKafkaCircuitBreaker",
////		MaxRequests: 1,
////		Interval:    60 * time.Second,
////		Timeout:     30 * time.Second,
////		ReadyToTrip: func(counts gobreaker.Counts) bool {
////			return counts.ConsecutiveFailures > 10
////		},
////		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
////			log.Printf("Circuit breaker state changed from %s to %s\n", from, to)
////		},
////	}
////	return gobreaker.NewCircuitBreaker(cbSettings)
////}
//
//func (messageService KafkaMessageService) SendMessage(request httprequest.PublishMessageRequest) error {
//	message := kafka.Message{
//		TopicPartition: kafka.TopicPartition{Topic: &request.Target, Partition: kafka.PartitionAny},
//		Value:          []byte(request.Message),
//		TimestampType:  kafka.TimestampLogAppendTime,
//	}
//
//	if messageService.circuitBreakerEnabled {
//		return messageService.sendMessageWithFallBack(message)
//	} else {
//		return messageService._sendMessage(messageService.primaryProducer, message)
//	}
//}
//
//func (messageService KafkaMessageService) SendMessageAsync(request httprequest.PublishMessageRequest) error {
//	return messageService.SendMessage(request)
//}
//
//func (messageService KafkaMessageService) _sendMessage(producer *kafka.Producer, message kafka.Message) error {
//	err := messageService.primaryProducer.Produce(&message, nil)
//	if err != nil {
//		log.Error(err)
//	}
//	return err
//}
//
//func (messageService KafkaMessageService) sendMessageWithFallBack(message kafka.Message) error {
//	cb := messageService.circuitBreaker
//
//	_, pErr := cb.Execute(func() (interface{}, error) {
//		// Attempt to send message to primary cluster
//		if err := messageService._sendMessage(messageService.primaryProducer, message); err != nil {
//			return nil, err
//		}
//		return nil, nil
//	})
//
//	if pErr != nil {
//		log.Error("Failed to send to primary cluster: %v\n", pErr)
//		if fErr := messageService._sendMessage(messageService.fallbackProducer, message); fErr != nil {
//			log.Error("Failed to send to fallback cluster: %v\n", fErr)
//		}
//		return errors.New("unable to send message")
//	}
//
//	return nil
//}
