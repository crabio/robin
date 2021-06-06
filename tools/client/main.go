package main

import (
	// External

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"

	// Internal
	"github.com/iakrevetkho/robin/internal/config"
)

func main() {
	// Load app configuration
	config, err := config.LoadConfig("../../config.yml")
	if err != nil {
		log.Fatalf("Couldn't load app config. %v", err)
	}
	log.Infof("Loaded config: %+v", config)

	// Connect to a server
	nc, _ := nats.Connect(config.NATS.Hostname)
	defer nc.Close()

	sendSuccessRequest(config, nc)
	sendFailedRequest(config, nc)
}
