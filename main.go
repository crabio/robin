package main

import (
	// External
	"time"

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"

	// Internal
	"github.com/iakrevetkho/robin/internal/config"
)

//Execution starts from main function
func main() {
	// Load app configuration
	config, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Couldn't load app config. %v", err)
	}
	log.Infof("Loaded config: %+v", config)

	// [begin request_reply]
	nc, err := nats.Connect(config.NATS.Hostname)
	if err != nil {
		log.Fatalf("Couldn't connect to NATS. %v", err)
	}
	defer nc.Close()

	// Send the request
	msg, err := nc.Request("time", nil, time.Second)
	if err != nil {
		log.Fatalf("Couldn't send NATS request. %v", err)
	}

	// Use the response
	log.Printf("Reply: %s", msg.Data)

	// Close the connection
	nc.Close()
	// [end request_reply]
}
