package main

import (
	// External

	log "github.com/sirupsen/logrus"

	// Internal
	apiconnectors "github.com/iakrevetkho/robin/internal/api/connectors"
	"github.com/iakrevetkho/robin/internal/config"
)

func main() {
	// Load app configuration
	config, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Couldn't load app config. %v", err)
	}
	log.Infof("Loaded config: %+v", config)

	// Init NATS connector
	conn, err := apiconnectors.Init(config)
	if err != nil {
		log.Fatalf("Couldn't connect to NATS. %v", err)
	}
	defer conn.Close()
}
