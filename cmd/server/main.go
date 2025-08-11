package main

import (
	"flag"
	"github.com/faanross/akkeDNS/internals/config"
	"github.com/faanross/akkeDNS/internals/server/server_https"
	"log"
	"os"
	"os/signal"
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

	// Create server config
	serverCfg := config.Config{
		Protocol:   cfg.Protocol,
		ServerAddr: cfg.ServerAddr,
		TlsCert:    cfg.TlsCert,
		TlsKey:     cfg.TlsKey,
	}

	// Create server
	server := server_https.NewHTTPSServer(&serverCfg)

	// Start the server in own goroutine
	go func() {
		log.Printf("Starting %s server on %s", cfg.Protocol, cfg.ServerAddr)
		if err := server.Start(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	// Graceful shutdown
	log.Println("Shutting down server...")
	if err := server.Stop(); err != nil {
		log.Printf("Error stopping server: %v", err)
	}

}
