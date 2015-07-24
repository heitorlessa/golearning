package main

type room struct {

	// Channel 'Forward' will be holding incoming messages
	forward chan []byte

	// Clients wishing to join the room
	join chan []byte

	// clients wishing to leave the room
	leave chan []byte

	// holds all clients in this room
	clients map[*client]bool
}
