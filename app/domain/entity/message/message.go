package message

type Message interface {
	GetMessage() string
	GetMessageTimeStamp() string
	GetMessageType() string
}
