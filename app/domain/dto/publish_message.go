package httprequest

import "shopdeck.com/messaging_service/app/constant/messageType"

type MessageMetadata struct {
}

type PublishMessageRequest struct {
	//MessageMetadata MessageMetadata `json:"metadata" binding:"required"`
	Timestamp   string                  `json:"timestamp" binding:"required"`
	ServiceId   string                  `json:"service_id" binding:"required"`
	MessageType messageType.MessageType `json:"message_type" binding:"required"`
	Target      string                  `json:"queue" binding:"required"`
	Message     string                  `json:"message" binding:"required"`
	//KafkaTopicDetails KafkaTopicDetails `json:"kafka" binding:""
}

type PublishMessageResponse struct {
	RequestId string `json:"request_id"`
}
