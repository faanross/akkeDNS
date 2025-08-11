package main

import (
	"context"
	"flag"
	"github.com/faanross/akkeDNS/internals/config"
	"github.com/faanross/akkeDNS/internals/models"
	"github.com/faanross/akkeDNS/internals/runloop"
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

	// NEW LOGIC HERE FOLLOWING RUNLOOP INTEGRATION

	// Create context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start run loop in goroutine
	go func() {
		log.Printf("Starting %s client run loop", cfg.Protocol)
		log.Printf("Delay: %v, Jitter: %d%%", cfg.Timing.Delay.Duration, cfg.Timing.Jitter)

		if err := runloop.RunLoopHTTPS(ctx, comm, cfg.Timing.Delay.Duration, cfg.Timing.Jitter); err != nil {
			log.Printf("Run loop error: %v", err)
		}
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Shutting down client...")
	cancel() // This will cause the run loop to exit
}
