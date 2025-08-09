package models

import "context"

// Message represents the data exchanged between client and server
type Message struct {
	Data []byte
}

// AgentCommunicator defines the contract for any communication protocol
type AgentCommunicator interface {
	// Send sends a message and waits for a response
	Send(ctx context.Context, msg Message) (Message, error)

	// Close gracefully shuts down the communicator
	Close() error
}

// Handler processes incoming messages
type Handler func(ctx context.Context, msg Message) (Message, error)

// Server defines the contract for protocol servers
type Server interface {
	// Start begins listening for requests
	Start(ctx context.Context, handler Handler) error

	// Stop gracefully shuts down the server
	Stop() error
}
