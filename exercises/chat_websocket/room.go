package main

type room struct {

	// Channel 'Forward' will be holding incoming messages
	forward chan []byte

	// Clients wishing to join the room
	join chan *client

	// clients wishing to leave the room
	leave chan *client

	// holds all clients in this room
	clients map[*client]bool
}

func (r *room) run() {
	for {
		// select ensures that code run one block of code at a time
		select {

		case client := <-r.join:
			// joining
			r.clients[client] = true

		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)

		// forward message to all clients if a message is ever sent to 'forward' channel
		case msg := <-r.forward:
			// goes through all clients (i :=0; i < # of clients; i++ -- short version)
			for client := range r.clients {
				select {

				// sends message on 'send' channel on each client as it will be picked up by 'write' method
				// in which will write down the message to browser's web socket
				case client.send <- msg:
					// send message

				// in case client doesnt accept the message, removes the client and tide things up
				default:
					// failed to send message
					delete(r.clients, client)
					close(client.send)
				}
			}

		}
	}

}
