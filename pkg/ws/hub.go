package ws

import (
	"example/model"
)

type (
	Hub interface {
		run()
	}

	hub struct {
		clients map[*Client]bool

		broadcast chan *model.Chat

		register chan *Client

		unregister chan *Client
	}
)

func NewHub() Hub {
	return &hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan *model.Chat),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			h.clients[client] = false
		case message := <-h.broadcast:
			for client := range h.clients {
				if client.user == message.To || client.user == message.From {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}
