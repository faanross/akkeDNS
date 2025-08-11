package runloop

import (
	"context"
	"encoding/json"
	"github.com/faanross/akkeDNS/internals/models"
	"github.com/faanross/akkeDNS/internals/server/server_https"
	"github.com/faanross/akkeDNS/internals/types"
	"github.com/faanross/akkeDNS/internals/utils"
	"log"
	"time"
)

func RunLoopHTTPS(ctx context.Context, comm models.AgentCommunicator, delay time.Duration, jitter int) error {

	for {
		// Check if context is cancelled
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// Send request
		msg := types.Message{
			Data: []byte("GET request"),
		}

		response, err := comm.Send(ctx, msg)
		if err != nil {
			log.Printf("Error sending request: %v", err)
			return err
		}

		// Parse and display response
		var httpsResp server_https.HTTPSResponse
		if err := json.Unmarshal(response.Data, &httpsResp); err != nil {
			log.Fatalf("Failed to parse response: %v", err)
		}

		log.Printf("Received response: change=%v", httpsResp.Change)

		// Calculate sleep duration with jitter
		sleepDuration := utils.CalculateSleepDuration(delay, jitter)
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
