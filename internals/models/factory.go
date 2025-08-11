package models

import (
	"fmt"
	"github.com/faanross/akkeDNS/internals/config"
	"github.com/faanross/akkeDNS/internals/server/server_https"
)

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

// NewServer creates a new server based on the protocol
func NewServer(cfg config.Config) (Server, error) {
	switch cfg.Protocol {
	case "https":
		return server_https.NewHTTPSServer(&cfg), nil
	case "dns":
		return nil, fmt.Errorf("DNS server not yet implemented")
	default:
		return nil, fmt.Errorf("unsupported protocol: %v", cfg.Protocol)
	}
}
