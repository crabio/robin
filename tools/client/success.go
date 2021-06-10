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

func sendSuccessRequest(config config.Config, nc *nats.Conn) {
	log.Info("Send success auth request")

	// Create test message
	msg := resources.Msg{
		Uuid: &resources.UUID{
			Value: uuid.NewV4().Bytes(),
		},
		Ts: &resources.Timestamp{
			Value: uint64(time.Now().Unix()),
		},
		// Payload: &resources.Msg_AuthUserRequest{
		// 	AuthUserRequest: &resources.AuthUserRequest{
		// 		Username: "test",
		// 		Password: "test",
		// 	},
		// },
	}

	// Serialize message
	msgBytes, err := proto.Marshal(&msg)
	if err != nil {
		log.Fatalf("Couldn't serialize msg. %v", err)
	}

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
