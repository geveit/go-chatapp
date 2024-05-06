package server

import (
	"encoding/json"
	"log"
	"sync"
)

type Hub struct {
	sync.RWMutex

	clients map[*client]bool

	broadcast  chan *Message
	register   chan *client
	unregister chan *client

	messages []*Message
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*client]bool),
		broadcast:  make(chan *Message),
		register:   make(chan *client),
		unregister: make(chan *client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.Lock()
			h.clients[client] = true
			h.Unlock()

			// Send all messages to the new client
			for _, msg := range h.messages {
				log.Printf("Sending message to new client: %v", msg)
				data, _ := json.Marshal(msg)
				client.send <- data
			}
		case client := <-h.unregister:
			h.Lock()
			if _, ok := h.clients[client]; ok {
				close(client.send)
				delete(h.clients, client)
			}
			h.Unlock()
		case msg := <-h.broadcast:
			h.RLock()
			h.messages = append(h.messages, msg)

			for client := range h.clients {
				data, _ := json.Marshal(msg)
				select {
				case client.send <- data:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			h.RUnlock()
		}
	}
}
