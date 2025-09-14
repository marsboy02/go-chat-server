package message

import (
	"encoding/json"
	"time"
)

// MessageType represents the type of message
type MessageType string

const (
	// MessageTypeChat represents a chat message
	MessageTypeChat MessageType = "chat"
	// MessageTypeJoin represents a user joining
	MessageTypeJoin MessageType = "join"
	// MessageTypeLeave represents a user leaving
	MessageTypeLeave MessageType = "leave"
	// MessageTypeError represents an error message
	MessageTypeError MessageType = "error"
)

// Message represents a chat message
type Message struct {
	Type      MessageType `json:"type"`
	Content   string      `json:"content"`
	Username  string      `json:"username"`
	Timestamp time.Time   `json:"timestamp"`
}

// NewChatMessage creates a new chat message
func NewChatMessage(username, content string) *Message {
	return &Message{
		Type:      MessageTypeChat,
		Content:   content,
		Username:  username,
		Timestamp: time.Now(),
	}
}

// NewJoinMessage creates a new join message
func NewJoinMessage(username string) *Message {
	return &Message{
		Type:      MessageTypeJoin,
		Content:   username + " joined the chat",
		Username:  username,
		Timestamp: time.Now(),
	}
}

// NewLeaveMessage creates a new leave message
func NewLeaveMessage(username string) *Message {
	return &Message{
		Type:      MessageTypeLeave,
		Content:   username + " left the chat",
		Username:  username,
		Timestamp: time.Now(),
	}
}

// NewErrorMessage creates a new error message
func NewErrorMessage(content string) *Message {
	return &Message{
		Type:      MessageTypeError,
		Content:   content,
		Username:  "System",
		Timestamp: time.Now(),
	}
}

// ToJSON converts the message to JSON bytes
func (m *Message) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// FromJSON creates a message from JSON bytes
func FromJSON(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}