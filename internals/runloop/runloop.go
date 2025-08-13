package runloop

import (
	"context"
	"encoding/json"
	"github.com/faanross/akkeDNS/internals/config"
	"github.com/faanross/akkeDNS/internals/models"
	"github.com/faanross/akkeDNS/internals/server/server_https"
	"github.com/faanross/akkeDNS/internals/utils"
	"log"
	"time"
)

func RunLoop(ctx context.Context, comm models.Agent, cfg *config.Config) error {
	currentProtocol := cfg.Protocol // Track which protocol we're using
	currentAgent := comm            // Track current agent (can change!)

	for {
		// Check if context is cancelled
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		response, err := currentAgent.Send(ctx)

		if err != nil {
			log.Printf("Error sending request: %v", err)
			// Don't exit - just sleep and try again
			time.Sleep(cfg.Timing.Delay)
			continue // Skip to next iteration
		}

		// BASED ON PROTOCOL, HANDLE PARSING DIFFERENTLY
		// Check if this is a transition signal
		if detectTransition(currentProtocol, response) {
			log.Printf("TRANSITION SIGNAL DETECTED! Switching protocols...")

			// Figure out what protocol to switch TO
			newProtocol := "dns"
			if currentProtocol == "dns" {
				newProtocol = "https"
			}

			// Create config for new protocol
			tempConfig := *cfg // Copy the config
			tempConfig.Protocol = newProtocol

			// Try to create new agent
			newAgent, err := models.NewAgent(&tempConfig)
			if err != nil {
				log.Printf("Failed to create %s agent: %v", newProtocol, err)
				// Don't switch if we can't create agent
			} else {
				// Update our tracking variables
				log.Printf("Successfully switched from %s to %s", currentProtocol, newProtocol)
				currentProtocol = newProtocol
				currentAgent = newAgent
			}
		} else {
			// Normal response - parse and log as before
			switch currentProtocol { // Note: use currentProtocol, not cfg.Protocol
			case "https":
				var httpsResp server_https.HTTPSResponse
				json.Unmarshal(response, &httpsResp)
				log.Printf("Received response: change=%v", httpsResp.Change)
			case "dns":
				ipAddr := string(response)
				log.Printf("Received response: IP=%v", ipAddr)
			}
		}

		// Calculate sleep duration with jitter
		sleepDuration := utils.CalculateSleepDuration(time.Duration(cfg.Timing.Delay), cfg.Timing.Jitter)
		log.Printf("Sleeping for %v", sleepDuration)

		// Sleep with cancellation support
		select {
		case <-time.After(sleepDuration):
			// Continue to next iteration
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// detectTransition checks if the response indicates we should switch protocols
func detectTransition(protocol string, response []byte) bool {
	switch protocol {
	case "https":
		var httpsResp server_https.HTTPSResponse
		if err := json.Unmarshal(response, &httpsResp); err != nil {
			return false
		}
		return httpsResp.Change

	case "dns":
		ipAddr := string(response)
		return ipAddr == "69.69.69.69"
	}

	return false
}
