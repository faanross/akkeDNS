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

	for {
		// Check if context is cancelled
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		response, err := comm.Send(ctx)
		if err != nil {
			log.Printf("Error sending request: %v", err)
			return err
		}

		// BASED ON PROTOCOL, HANDLE PARSING DIFFERENTLY

		switch cfg.Protocol {
		case "https":
			// Parse and display response
			var httpsResp server_https.HTTPSResponse
			if err := json.Unmarshal(response, &httpsResp); err != nil {
				log.Fatalf("Failed to parse response: %v", err)
			}

			log.Printf("Received response: change=%v", httpsResp.Change)
		case "dns":
			ipAddr := string(response)
			log.Printf("Received response: IP=%v", ipAddr)

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
