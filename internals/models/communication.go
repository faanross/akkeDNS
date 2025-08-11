package models

import (
	"context"
	"github.com/faanross/akkeDNS/internals/types"
)

// AgentCommunicator defines the contract for any communication protocol
type AgentCommunicator interface {
	// Send sends a message and waits for a response
	Send(ctx context.Context, msg types.Message) (types.Message, error)

	// Close gracefully shuts down the communicator
	Close() error
}

// Server defines the contract for protocol servers
type Server interface {
	// Start begins listening for requests
	Start() error

	// Stop gracefully shuts down the server
	Stop() error
}
