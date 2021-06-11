package main

import (
	// External
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"

	// Internal
	"github.com/iakrevetkho/robin/internal/config"
	resources "github.com/iakrevetkho/robin/internal/resources"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

func LoginHandler(config config.Config, nc *nats.Conn, c *gin.Context) {
	log.Info("Process login request")

	// Create test message
	msg := resources.Msg{
		Uuid: &resources.UUID{
			Value: uuid.NewV4().Bytes(),
		},
		Ts: &resources.Timestamp{
			Value: uint64(time.Now().Unix()),
		},
		Payload: &resources.Msg_LoginRequest{
			LoginRequest: &resources.LoginRequest{
				Provider: resources.AuthProviderEnum_google,
			},
		},
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

	c.Redirect(http.StatusTemporaryRedirect, response.GetLoginResponse().GetUrl())
}
