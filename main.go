package main

import (
	// External

	"time"

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"

	// Internal
	apicontrollers "github.com/iakrevetkho/robin/internal/api/controllers"
	"github.com/iakrevetkho/robin/internal/config"
)

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

	// Subscribe on queue in the subject
	sub, err := nc.QueueSubscribe(config.NATS.Request.Subject, config.NATS.Request.Queue, apicontrollers.ProcessNatsMsg)
	if err != nil {
		log.Fatalf("Couldn't subscribe on request topic '%s' with queue '%s'. %v", config.NATS.Request.Subject, config.NATS.Request.Queue, err)
	}

	// Send the request
	msg, err := nc.Request(config.NATS.Request.Subject, []byte("lawndla"), time.Second)
	if err != nil {
		log.Fatalf("Couldn't send NATS request. %v", err)
	}

	// Use the response
	log.Printf("Reply: %s", msg.Data)

	// Close subscription
	err = sub.Unsubscribe()
	if err != nil {
		log.Errorf("Couldn't unsubscribe from request topic '%s' with queue '%s'. %v", config.NATS.Request.Subject, config.NATS.Request.Queue, err)
	}

	// Close the connection
	nc.Close()
	// [end request_reply]
}
