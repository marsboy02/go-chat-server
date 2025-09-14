package hub

import (
	"log"
	"sync"

	"go-chat-server/internal/message"
	"go-chat-server/internal/types"
)

// Hub maintains the set of active clients and broadcasts messages to the clients
type Hub struct {
	// Registered clients
	clients map[types.HubClient]bool

	// Inbound messages from the clients
	broadcast chan *message.Message

	// Register requests from the clients
	register chan types.HubClient

	// Unregister requests from clients
	unregister chan types.HubClient

	// Mutex to protect concurrent access to clients map
	mutex sync.RWMutex
}

// New creates a new Hub
func New() *Hub {
	return &Hub{
		clients:    make(map[types.HubClient]bool),
		broadcast:  make(chan *message.Message),
		register:   make(chan types.HubClient),
		unregister: make(chan types.HubClient),
	}
}

// Run starts the hub's main loop
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mutex.Lock()
			h.clients[client] = true
			h.mutex.Unlock()

			// Send join message to all clients
			joinMsg := message.NewJoinMessage(client.GetUsername())
			h.broadcastToAll(joinMsg)

			log.Printf("Client %s joined. Total clients: %d", client.GetUsername(), len(h.clients))

		case client := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				h.mutex.Unlock()

				// Send leave message to all clients
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

// Register adds a client to the hub
func (h *Hub) Register(client types.HubClient) {
	h.register <- client
}

// Unregister removes a client from the hub
func (h *Hub) Unregister(client types.HubClient) {
	h.unregister <- client
}

// Broadcast sends a message to all connected clients
func (h *Hub) Broadcast(msg *message.Message) {
	h.broadcast <- msg
}

// broadcastToAll sends a message to all connected clients
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

// GetClientCount returns the number of connected clients
func (h *Hub) GetClientCount() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return len(h.clients)
}

// GetConnectedUsers returns a slice of usernames of all connected users
func (h *Hub) GetConnectedUsers() []string {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	users := make([]string, 0, len(h.clients))
	for client := range h.clients {
		users = append(users, client.GetUsername())
	}
	return users
}