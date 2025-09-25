package handler

import (
	"log"
	"net/http"

	"go-chat-server/internal/client"
	"go-chat-server/internal/hub"
)

type WebSocketHandler struct {
	hub *hub.Hub
}

func NewWebSocketHandler(hub *hub.Hub) *WebSocketHandler {
	return &WebSocketHandler{
		hub: hub,
	}
}

func (h *WebSocketHandler) ServeWS(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		username = "Anonymous"
	}

	conn, err := client.Upgrade(w, r)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	c := client.NewClient(h.hub, conn, username)

	h.hub.Register(c)

	c.Run()
}