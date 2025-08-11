package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/faanross/akkeDNS/internals/config"
	"github.com/faanross/akkeDNS/internals/models"
	"github.com/faanross/akkeDNS/internals/server/server_https"
	"github.com/faanross/akkeDNS/internals/types"
	"log"
)

const pathToYAML = "./configs/config.yaml"

func main() {
	// Command line flag for config file path
	configPath := flag.String("config", pathToYAML, "path to configuration file")
	flag.Parse()

	// Load configuration
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	commCfg := config.AgentConfig{
		Protocol:   cfg.Protocol,
		ServerAddr: cfg.ServerAddr,
		AgentAddr:  cfg.ClientAddr,
	}

	comm, err := models.NewCommunicator(commCfg)
	if err != nil {
		log.Fatalf("Failed to create communicator: %v", err)
	}
	defer comm.Close()

	// Send a test message
	msg := types.Message{
		Data: []byte("Hello from client"),
	}

	log.Printf("Sending request to %s server...", cfg.Protocol)
	response, err := comm.Send(context.Background(), msg)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Parse and display response
	var httpsResp server_https.HTTPSResponse
	if err := json.Unmarshal(response.Data, &httpsResp); err != nil {
		log.Fatalf("Failed to parse response: %v", err)
	}

	log.Printf("Received response: change=%v", httpsResp.Change)

}
