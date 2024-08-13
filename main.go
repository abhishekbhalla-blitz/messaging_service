package main

import (
	"shopdeck.com/messaging_service/app/router"
	"shopdeck.com/messaging_service/config"
	"shopdeck.com/messaging_service/initialize"
	"strconv"
)

//func produceConfluentKafkaGo(numMessages int, topic *string, value []byte) {
//
//	// ~380k/s
//
//	var p, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "35.200.142.121:9092", "linger.ms": 100, "batch.num.messages": 50})
//	if err != nil {
//		log.Printf("could not set up kafka producer: %s", err.Error())
//		os.Exit(1)
//	}
//
//	done := make(chan bool)
//	// go func() {
//	// 	var msgCount int
//	// 	for e := range p.Events() {
//	// 		msg := e.(*kafka.Message)
//	// 		if msg.TopicPartition.Error != nil {
//	// 			log.Printf("delivery report error: %v", msg.TopicPartition.Error)
//	// 			os.Exit(1)
//	// 		}
//	// 		msgCount++
//	// 		if msgCount >= numMessages {
//	// 			done <- true
//	// 		}
//	// 	}
//	// }()
//
//	// defer p.Close()
//
//	var start = time.Now()
//	for j := 0; j < numMessages; j++ {
//		p.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: topic}, Value: value}
//	}
//	<-done
//	elapsed := time.Since(start)
//
//	log.Printf("[confluent-kafka-go producer] msg/s: %f", (float64(numMessages) / elapsed.Seconds()))
//}

func init() {
	config.InitConfig()
	initialize.InitLog()
}

func main() {
	configuration := config.GetConfiguration()
	init := initialize.Init()
	app := router.Init(init)
	app.Run(":" + strconv.Itoa(configuration.Server.Port))
}
