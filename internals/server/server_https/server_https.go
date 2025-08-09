package server_https

import "net/http"

// HTTPSServer implements the Server interface for HTTPS
type HTTPSServer struct {
	addr   string
	server *http.Server
}

// HTTPSResponse represents the JSON response for HTTPS
type HTTPSResponse struct {
	Change bool `json:"change"`
}

// NewHTTPSServer creates a new HTTPS server
func NewHTTPSServer(addr string) *HTTPSServer {
	return &HTTPSServer{
		addr: addr,
	}
}
