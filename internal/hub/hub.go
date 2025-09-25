package hub

import (
	"log"
	"sync"

	"go-chat-server/internal/message"
	"go-chat-server/internal/types"
)

type Hub struct {
	clients map[types.HubClient]bool

	broadcast chan *message.Message

	register chan types.HubClient

	unregister chan types.HubClient

	mutex sync.RWMutex
}

func New() *Hub {
	return &Hub{
		clients:    make(map[types.HubClient]bool),
		broadcast:  make(chan *message.Message),
		register:   make(chan types.HubClient),
		unregister: make(chan types.HubClient),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mutex.Lock()
			h.clients[client] = true
			h.mutex.Unlock()

			joinMsg := message.NewJoinMessage(client.GetUsername())
			h.broadcastToAll(joinMsg)

			log.Printf("Client %s joined. Total clients: %d", client.GetUsername(), len(h.clients))

		case client := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				h.mutex.Unlock()

				leaveMsg := message.NewLeaveMessage(client.GetUsername())
				h.broadcastToAll(leaveMsg)

				log.Printf("Client %s left. Total clients: %d", client.GetUsername(), len(h.clients))
			} else {
				h.mutex.Unlock()
			}

		case msg := <-h.broadcast:
			h.broadcastToAll(msg)
		}
	}
}

func (h *Hub) Register(client types.HubClient) {
	h.register <- client
}

func (h *Hub) Unregister(client types.HubClient) {
	h.unregister <- client
}

func (h *Hub) Broadcast(msg *message.Message) {
	h.broadcast <- msg
}

func (h *Hub) broadcastToAll(msg *message.Message) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	for client := range h.clients {
		select {
		default:
			client.Send(msg)
		}
	}
}

func (h *Hub) GetClientCount() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return len(h.clients)
}

func (h *Hub) GetConnectedUsers() []string {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	users := make([]string, 0, len(h.clients))
	for client := range h.clients {
		users = append(users, client.GetUsername())
	}
	return users
}