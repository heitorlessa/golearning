package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"trace"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

type room struct {

	// Channel 'Forward' will be holding incoming messages
	forward chan []byte

	// Clients wishing to join the room
	join chan *client

	// clients wishing to leave the room
	leave chan *client

	// holds all clients in this room
	clients map[*client]bool

	// tracer will receive trace information of activity
	// in the room
	tracer trace.Tracer
}

// helper to easily create new room (initialize map & channels)
func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  trace.Off(),
	}
}

func (r *room) run() {
	for {
		// select ensures that code run one block of code at a time
		select {

		case client := <-r.join:
			// joining
			r.clients[client] = true
			r.tracer.Trace("New client joined")

		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left")

		// forward message to all clients if a message is ever sent to 'forward' channel
		case msg := <-r.forward:
			// goes through all clients (i :=0; i < # of clients; i++ -- short version)
			for client := range r.clients {
				select {

				// sends message on 'send' channel on each client as it will be picked up by 'write' method
				// in which will write down the message to browser's web socket
				case client.send <- msg:
					// send message
					r.tracer.Trace(" -- sent to client")

				// in case client doesnt accept the message, removes the client and tide things up
				default:
					// failed to send message
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace(" -- failed to send, cleaned up client")
				}
			}
		}
	}
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// need to upgrade socket (HTTP Request) into webSocket
	socket, err := upgrader.Upgrade(w, req, nil)

	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	// creates a client that will join a channel later on
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r, // weird as it is this ',' is correct
	}

	// makes client join (in other words, sends message to 'join' channel)
	r.join <- client

	// makes client leave but defer it to ensure clients successfuly join first
	defer func() { r.leave <- client }() // similar to JS promises
	go client.write()                    // goroutine, in other words, run this portion of code in a new thread
	client.read()

}
