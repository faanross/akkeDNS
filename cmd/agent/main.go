package main

import "github.com/faanross/akkeDNS/internals/models"

func main() {
	agentConfig := models.AgentConfig{
		Protocol:   models.ProtocolDNS,
		ServerAddr: "127.0.0.1:6565",
		AgentAddr:  "127.0.0.1:0",
	}

	_, err := models.NewCommunicator(agentConfig)
	if err != nil {
		panic(err)
	}

}
