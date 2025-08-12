package agent_https

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/faanross/akkeDNS/internals/server/server_https"
	"io"
	"net/http"
)

// HTTPSAgent implements the Communicator interface for HTTPS
type HTTPSAgent struct {
	serverAddr string
	client     *http.Client
}

// NewHTTPSAgent creates a new HTTPS agent
func NewHTTPSAgent(serverAddr string) *HTTPSAgent {
	// Create TLS config that accepts self-signed certificates
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	// Create HTTP client with custom TLS config
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	return &HTTPSAgent{
		serverAddr: serverAddr,
		client:     client,
	}
}

// Send implements Communicator.Send for HTTPS
func (c *HTTPSAgent) Send(ctx context.Context) ([]byte, error) {
	// Construct the URL
	url := fmt.Sprintf("https://%s/", c.serverAddr)

	// Create GET request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// Send request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("server returned status %d: %s", resp.StatusCode, body)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	// Verify it's valid JSON by parsing it
	var httpsResp server_https.HTTPSResponse
	if err := json.Unmarshal(body, &httpsResp); err != nil {
		return nil, fmt.Errorf("invalid JSON response: %w", err)
	}

	// Return the raw JSON as message data
	return body, nil
}
