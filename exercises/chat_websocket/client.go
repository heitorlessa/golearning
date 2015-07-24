package main

import (
	"github.com/gorilla/websocket"
)

// client represents a single chatting user
type client struct {
	socket *websocket.Conn //
	send   chan []byte     // go channel
	room   *room
}

// read method that reads from message (ReadMessage through websocket) from 'forward' chan
func (c *client) read() {
	for {

		if _, msg, err := c.socket.ReadMessage(); err == nil {
			// read from channel '<-' (in-memory message queue [thread-safe])
			c.room.forward <- msg
		} else {
			break
		}
	}

	c.socket.Close()
}

// write method that accepts messages from 'send' chan and writes (WriteMessage 'Text' message through websocket)
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
