package message

import (
	constant "shopdeck.com/messaging_service/app/constant/messageType"
)

type KafkaMessage struct {
	message   string
	timestamp string
	topic     string
	partition string
}

func (message *KafkaMessage) GetMessage() string {
	return message.message
}

func (message *KafkaMessage) GetMessageType() constant.MessageType {
	return constant.KAFKA
}

func (message *KafkaMessage) GetMessageTimeStamp() string {
	return message.timestamp
}

func (message *KafkaMessage) GetTopic() string {
	return message.topic
}

func (message *KafkaMessage) GetPartition() string {
	return message.partition
}
