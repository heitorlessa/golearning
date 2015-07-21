package main

type room struct {
	
	// Channel 'Forward' will be holding incoming messages
	forward chan []byte
}