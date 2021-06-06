package main

import (
	// External
	"time"

	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"

	// Internal
	"github.com/iakrevetkho/robin/internal/config"
	resources "github.com/iakrevetkho/robin/internal/resources"
)

func main() {
	// Load app configuration
	config, err := config.LoadConfig("../../config.yml")
	if err != nil {
		log.Fatalf("Couldn't load app config. %v", err)
	}
	log.Infof("Loaded config: %+v", config)

	// Create test message
	msg := resources.Msg{
		Uuid: &resources.UUID{
			Value: uuid.NewV4().Bytes(),
		},
		Ts: &resources.Timestamp{
			Value: uint64(time.Now().Unix()),
		},
		Payload: &resources.Msg_AuthUserRequest{
			AuthUserRequest: &resources.AuthUserRequest{
				Username: "test",
				Password: "test",
			},
		},
	}

	// Serialize message
	msgBytes, err := proto.Marshal(&msg)
	if err != nil {
		log.Fatalf("Couldn't serialize msg. %v", err)
	}

	// Connect to a server
	nc, _ := nats.Connect(config.NATS.Hostname)
	defer nc.Close()

	// Send request
	responseProto, err := nc.Request(config.NATS.Request.Subject, msgBytes, 1*time.Second)
	if err != nil {
		log.Fatalf("Couldn't send request. %v", err)
	}

	// Parse response
	response := resources.Msg{}
	err = proto.Unmarshal(responseProto.Data, &response)
	if err != nil {
		log.Fatalf("Couldn't deserialize response. %v", err)
	}

	log.Infof("Response: %+v", response.GetPayload())
}
