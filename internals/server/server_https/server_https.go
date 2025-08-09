package server_https

import (
	"context"
	"github.com/faanross/akkeDNS/internals/config"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

// HTTPSServer implements the Server interface for HTTPS
type HTTPSServer struct {
	addr    string
	server  *http.Server
	tlsCert string
	tlsKey  string
}

// HTTPSResponse represents the JSON response for HTTPS
type HTTPSResponse struct {
	Change bool `json:"change"`
}

// NewHTTPSServer creates a new HTTPS server
func NewHTTPSServer(cfg *config.Config) *HTTPSServer {
	return &HTTPSServer{
		addr:    cfg.ServerAddr,
		tlsCert: cfg.TlsCert,
		tlsKey:  cfg.TlsKey,
	}
}

// Start implements Server.Start for HTTPS
func (s *HTTPSServer) Start(ctx context.Context) error {
	// Create Chi router
	r := chi.NewRouter()

	// Define our GET endpoint
	r.Get("/", RootHandler)

	// Create the HTTP server
	s.server = &http.Server{
		Addr:    s.addr,
		Handler: r,
	}

	// Start the server
	return s.server.ListenAndServeTLS(s.tlsCert, s.tlsKey)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Endpoint %s has been hit by agent\n", r.URL.Path)

	// This is the handler on the endpoint

	// The agent hits the endpoint, it then responds with a struct HTTPSResponse, with change = false
	// Later I will add logic that may change this to true, but for now don't worry about it, just make it default false

	// That's all that happens, here we create struct HTTPSResponse, change to false, serialize, send to client

}
