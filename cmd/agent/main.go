package main

import (
	"flag"
	"github.com/faanross/akkeDNS/internals/config"
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

	log.Printf("Loaded configuration:\n")
	log.Printf("-> Client: %s\n", cfg.ClientAddr)
	log.Printf("-> Server: %s\n", cfg.ServerAddr)
	log.Printf("-> Delay: %v\n", cfg.Timing.Delay.Duration)
	log.Printf("-> Jitter: %d%%\n", cfg.Timing.Jitter)
	log.Printf("-> Starting Protocol: %s\n", cfg.Protocol)
	log.Printf("-> Cert location: %s\n", cfg.TlsCert)
	log.Printf("-> Key location: %s\n", cfg.TlsKey)
}
