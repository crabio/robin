package main

import (
	// External
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"

	// Internal
	"github.com/iakrevetkho/robin/internal/config"
	proto_resources "github.com/iakrevetkho/robin/internal/proto_resources"
)

func sendFailedRequest(config config.Config, nc *nats.Conn) error {
	log.Info("Send failed auth request")

	// Send request
	responseProto, err := nc.Request(config.NATS.Request.Subject, []byte("blabla"), 1*time.Second)
	if err != nil {
		return fmt.Errorf("Couldn't send request. %v", err)
	}

	// Parse response
	response := proto_resources.Msg{}
	err = proto.Unmarshal(responseProto.Data, &response)
	if err != nil {
		return fmt.Errorf("Couldn't deserialize response. %v", err)
	}

	log.Infof("Response: %+v", response.GetPayload())

	return nil
}
