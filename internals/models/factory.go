package models

import (
	"fmt"
	"github.com/faanross/akkeDNS/internals/agent/agent_dns"
	"github.com/faanross/akkeDNS/internals/agent/agent_https"
	"github.com/faanross/akkeDNS/internals/config"
	"github.com/faanross/akkeDNS/internals/server/server_dns"
	"github.com/faanross/akkeDNS/internals/server/server_https"
)

// NewAgent creates a new communicator based on the protocol
func NewAgent(cfg *config.Config) (Agent, error) {
	switch cfg.Protocol {
	case "https":
		return agent_https.NewHTTPSAgent(cfg.ServerAddr), nil
	case "dns":
		return agent_dns.NewDNSAgent(cfg.ServerAddr), nil
	default:
		return nil, fmt.Errorf("unsupported protocol: %v", cfg.Protocol)
	}
}

// NewServer creates a new server based on the protocol
func NewServer(cfg *config.Config) (Server, error) {
	switch cfg.Protocol {
	case "https":
		return server_https.NewHTTPSServer(cfg), nil
	case "dns":
		return server_dns.NewDNSServer(cfg), nil
	default:
		return nil, fmt.Errorf("unsupported protocol: %v", cfg.Protocol)
	}
}
