package models

import (
	"fmt"
	"github.com/faanross/akkeDNS/internals/agent/agent_https"
	"github.com/faanross/akkeDNS/internals/config"
	"github.com/faanross/akkeDNS/internals/server/server_dns"
	"github.com/faanross/akkeDNS/internals/server/server_https"
)

// NewCommunicator creates a new communicator based on the protocol
func NewCommunicator(cfg config.AgentConfig) (AgentCommunicator, error) {
	switch cfg.Protocol {
	case "https":
		return agent_https.NewHHTTPSAgent(cfg.ServerAddr), nil
	case "dns":
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
		return server_dns.NewDNSServer(&cfg), nil
	default:
		return nil, fmt.Errorf("unsupported protocol: %v", cfg.Protocol)
	}
}
