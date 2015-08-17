package main

import (
	"github.com/gorilla/websocket"
)

// client represents a single chatting user
type client struct {
	// socket is the web socket for this client
	socket *websocket.Conn
	// send is the channel on which messages are sent
	send chan *message
	// room is the room this client is chatting in
	room *room
	// userData holds information about the user
	userData map[string]interface{}
}

// read method that reads from message (ReadMessage through websocket) from 'forward' chan
func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.Name = c.userData["name"].(string)
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
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
