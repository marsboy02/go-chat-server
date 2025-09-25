package message

import (
	"encoding/json"
	"time"
)

type MessageType string

const (
	MessageTypeChat MessageType = "chat"
	MessageTypeJoin MessageType = "join"
	MessageTypeLeave MessageType = "leave"
	MessageTypeError MessageType = "error"
)

type Message struct {
	Type      MessageType `json:"type"`
	Content   string      `json:"content"`
	Username  string      `json:"username"`
	Timestamp time.Time   `json:"timestamp"`
}

func NewChatMessage(username, content string) *Message {
	return &Message{
		Type:      MessageTypeChat,
		Content:   content,
		Username:  username,
		Timestamp: time.Now(),
	}
}

func NewJoinMessage(username string) *Message {
	return &Message{
		Type:      MessageTypeJoin,
		Content:   username + " joined the chat",
		Username:  username,
		Timestamp: time.Now(),
	}
}

func NewLeaveMessage(username string) *Message {
	return &Message{
		Type:      MessageTypeLeave,
		Content:   username + " left the chat",
		Username:  username,
		Timestamp: time.Now(),
	}
}

func NewErrorMessage(content string) *Message {
	return &Message{
		Type:      MessageTypeError,
		Content:   content,
		Username:  "System",
		Timestamp: time.Now(),
	}
}

func (m *Message) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

func FromJSON(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}