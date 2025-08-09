package models

import "fmt"

// AgentConfig holds configuration for creating an Agent Communicator
type AgentConfig struct {
	Protocol   Protocol
	ServerAddr string
	AgentAddr  string
}

// NewCommunicator creates a new communicator based on the protocol
func NewCommunicator(cfg AgentConfig) (AgentCommunicator, error) {
	switch cfg.Protocol {
	case ProtocolHTTPS:
		return nil, fmt.Errorf("HTTPS not yet implemented")
	case ProtocolDNS:
		return nil, fmt.Errorf("DNS not yet implemented")
	default:
		return nil, fmt.Errorf("unsupported protocol: %v", cfg.Protocol)
	}
}
