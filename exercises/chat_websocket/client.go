package main

import (
	"github.com/gorilla/websocket"
)

// client represents a single chatting user
type client struct {
	socket *websocket.Conn //
	send chan []byte // go channel
	room *room
}

func (c *client) read() {
	for {
		
		if _, msg, err := c.socket.ReadMessage(); err == nill {
			// read from channel '<-' (in-memory message queue [thread-safe])
			c.room.forward <- msg
		}

		else {
			break
		}

	}

	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}

	c.socket.Close()
}