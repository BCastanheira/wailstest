package backend

import(
	"context"
)

type Test struct {
	Name string `json:"name"`
	Type string `json:"type"`
	ctx  context.Context
}

type ChannelMessage struct {
	Type    string
	Payload interface{}
}

type MessageHandler interface {
	SendMessage(message ChannelMessage) error
}

type ChannelMessageHandler struct {
	MessageChannel chan ChannelMessage
}