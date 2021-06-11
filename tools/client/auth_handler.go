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

type AuthRequest struct {
	AuthCode string `form:"code" binding:"required"`
}

func AuthHandler(config config.Config, nc *nats.Conn, c *gin.Context) {
	log.Info("Process auth request")

	// Parse request body
	requestBody := AuthRequest{}
	err := c.BindQuery(&requestBody)
	if err != nil {
		c.String(400, "Couldn't parse auth request args. %v", err)
		c.Abort()
		return
	}

	// Create test message
	msg := resources.Msg{
		Uuid: &resources.UUID{
			Value: uuid.NewV4().Bytes(),
		},
		Ts: &resources.Timestamp{
			Value: uint64(time.Now().Unix()),
		},
		Payload: &resources.Msg_AuthRequest{
			AuthRequest: &resources.AuthRequest{
				Provider: resources.AuthProviderEnum_google,
				AuthCode: requestBody.AuthCode,
			},
		},
	}

	// Serialize message
	msgBytes, err := proto.Marshal(&msg)
	if err != nil {
		c.String(500, "Couldn't serialize msg. %v", err)
		c.Abort()
		return
	}

	// Send request
	responseProto, err := nc.Request(config.NATS.Request.Subject, msgBytes, 1*time.Second)
	if err != nil {
		c.String(500, "Couldn't send request. %v", err)
		c.Abort()
		return
	}

	// Parse response
	response := resources.Msg{}
	err = proto.Unmarshal(responseProto.Data, &response)
	if err != nil {
		c.String(500, "Couldn't deserialize response. %v", err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.String())
}
