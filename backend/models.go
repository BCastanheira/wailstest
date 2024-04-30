package backend

type Test struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	EventHandler *EventHandler
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
