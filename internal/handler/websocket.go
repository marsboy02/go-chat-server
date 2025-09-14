package handler

import (
	"log"
	"net/http"

	"go-chat-server/internal/client"
	"go-chat-server/internal/hub"
)

// WebSocketHandler handles WebSocket connections
type WebSocketHandler struct {
	hub *hub.Hub
}

// NewWebSocketHandler creates a new WebSocket handler
func NewWebSocketHandler(hub *hub.Hub) *WebSocketHandler {
	return &WebSocketHandler{
		hub: hub,
	}
}

// ServeWS handles websocket requests from the peer
func (h *WebSocketHandler) ServeWS(w http.ResponseWriter, r *http.Request) {
	// Get username from query parameter
	username := r.URL.Query().Get("username")
	if username == "" {
		username = "Anonymous"
	}

	// Upgrade connection to WebSocket
	conn, err := client.Upgrade(w, r)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	// Create new client
	c := client.NewClient(h.hub, conn, username)

	// Register client with hub (cast to HubClient interface)
	h.hub.Register(c)

	// Start client's read and write pumps
	c.Run()
}