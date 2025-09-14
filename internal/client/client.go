package client

import (
	"log"
	"net/http"
	"time"

	"go-chat-server/internal/message"
	"go-chat-server/internal/types"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin
	},
}


// Client represents a websocket client
type Client struct {
	// The websocket connection
	conn *websocket.Conn

	// Buffered channel of outbound messages
	send chan *message.Message

	// The hub that manages this client
	hub types.Hub

	// Client username
	Username string
}

// NewClient creates a new websocket client
func NewClient(hub types.Hub, conn *websocket.Conn, username string) *Client {
	return &Client{
		conn:     conn,
		send:     make(chan *message.Message, 256),
		hub:      hub,
		Username: username,
	}
}

// Upgrade upgrades HTTP connection to websocket
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return upgrader.Upgrade(w, r, nil)
}

// Run starts the client's read and write goroutines
func (c *Client) Run() {
	go c.writePump()
	go c.readPump()
}

// readPump pumps messages from the websocket connection to the hub
func (c *Client) readPump() {
	defer func() {
		c.hub.Unregister(c)
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, messageBytes, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("websocket error: %v", err)
			}
			break
		}

		// Parse the incoming message
		incomingMsg, err := message.FromJSON(messageBytes)
		if err != nil {
			log.Printf("error parsing message: %v", err)
			continue
		}

		// Create a new message with the client's username and current timestamp
		chatMsg := message.NewChatMessage(c.Username, incomingMsg.Content)
		c.hub.Broadcast(chatMsg)
	}
}

// writePump pumps messages from the hub to the websocket connection
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			msgBytes, err := msg.ToJSON()
			if err != nil {
				log.Printf("error marshaling message: %v", err)
				continue
			}
			w.Write(msgBytes)

			// Add queued chat messages to the current websocket message
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				queuedMsg := <-c.send
				queuedMsgBytes, err := queuedMsg.ToJSON()
				if err != nil {
					log.Printf("error marshaling queued message: %v", err)
					continue
				}
				w.Write(queuedMsgBytes)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// Send sends a message to the client
func (c *Client) Send(msg *message.Message) {
	select {
	case c.send <- msg:
	default:
		close(c.send)
	}
}

// GetUsername returns the client's username
func (c *Client) GetUsername() string {
	return c.Username
}