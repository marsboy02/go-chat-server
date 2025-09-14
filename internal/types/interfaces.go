package types

import "go-chat-server/internal/message"

// HubClient represents a client interface for the hub
type HubClient interface {
	Send(msg *message.Message)
	GetUsername() string
}

// Hub represents the hub that manages all clients
type Hub interface {
	Register(client HubClient)
	Unregister(client HubClient)
	Broadcast(message *message.Message)
}