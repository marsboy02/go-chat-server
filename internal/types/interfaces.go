package types

import "go-chat-server/internal/message"

type HubClient interface {
	Send(msg *message.Message)
	GetUsername() string
}

type Hub interface {
	Register(client HubClient)
	Unregister(client HubClient)
	Broadcast(message *message.Message)
}