// Adapted from fasthttp's websocket package (which was adapted from gorilla/websocket)

package ws

import (
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

// Hub maintains the set of active Clients and broadcasts messages to the
// Clients.
type Hub struct {
	// Registered Clients.
	Clients map[*Client]bool

	// Inbound messages from the Clients.
	Broadcast chan []byte

	// Register requests from the Clients.
	Register chan *Client

	// Unregister requests from Clients.
	Unregister chan *Client
	Log        *zap.SugaredLogger
}

func NewHub(log *zap.SugaredLogger) *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Log:        log,
	}
}

type Message struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (h *Hub) ActionBroadcast(message string) {
	body, err := jsoniter.Marshal(&Message{Type: "action", Message: message})
	if err != nil {
		h.Log.Errorw("failed to marshal logged message", "error", err)
		return
	}

	h.Broadcast <- body
}

func (h *Hub) LogBroadcast(message string) {
	body, err := jsoniter.Marshal(&Message{Type: "log", Message: message})
	if err != nil {
		h.Log.Errorw("failed to marshal logged message", "error", err)
		return
	}

	h.Broadcast <- body
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
