package messageType

import "shopdeck.com/messaging_service/app/constant"

type MessageType string

type messageTypeEnum struct {
	Name  string
	Value MessageType
}

const (
	UNDEFINED  MessageType = "UNDEFINED"
	KAFKA      MessageType = "KAFKA"
	GCP_PUBSUB MessageType = "GCP_PUBSUB"
	AWS_SQS    MessageType = "AWS_SQS"
)

var messageTypeEnums = []messageTypeEnum{
	{"UNDEFINED", UNDEFINED},
	{"KAFKA", KAFKA},
	{"GCP_PUBSUB", GCP_PUBSUB},
	{"AWS_SQS", AWS_SQS},
}

func (mt messageTypeEnum) GetName() string {
	return mt.Name
}

func (mt messageTypeEnum) GetValue() MessageType {
	return mt.Value
}

func (mt messageTypeEnum) FindByName(name string) (constant.Enum[MessageType], bool) {
	for _, enum := range messageTypeEnums {
		if enum.Name == name {
			return enum, true
		}
	}
	return nil, false
}

func (mt messageTypeEnum) FindByValue(value MessageType) (constant.Enum[MessageType], bool) {
	for _, enum := range messageTypeEnums {
		if enum.Value == value {
			return enum, true
		}
	}
	return nil, false
}
