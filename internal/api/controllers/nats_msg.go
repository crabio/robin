package apicontrollers

import (
	// External

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"

	// Internal
	helpers "github.com/iakrevetkho/robin/internal/api/helpers"
	apiresources "github.com/iakrevetkho/robin/internal/api/resources"
	apirouters "github.com/iakrevetkho/robin/internal/api/routers"
	resources "github.com/iakrevetkho/robin/internal/resources"
)

// Method for processing messages from NATS broker
func NatsMessage(controllerData apiresources.ControllerData, msg *nats.Msg) {
	// Check that msg is request
	if msg.Reply == "" {
		helpers.NatsErrorResponse(msg, &resources.UUID{}, "Await request, but receive msg without reply: %+v", msg)
		return
	}

	// Parse message protobuf
	protoMsg := resources.Msg{}
	if err := proto.Unmarshal(msg.Data, &protoMsg); err != nil {
		helpers.NatsErrorResponse(msg, &resources.UUID{}, "Failed to parse proto msg '%s': %v", msg.Data, err)
		return
	}

	// Go to msg router
	response, err := apirouters.RouteMsg(controllerData, &protoMsg)
	if err != nil {
		helpers.NatsErrorResponse(msg, &resources.UUID{}, "Couldn't process msg '%+v': %v", protoMsg.GetPayload(), err)
		return
	}

	// Serialize response
	responseProto, err := proto.Marshal(response)
	if err != nil {
		helpers.NatsErrorResponse(msg, &resources.UUID{}, "Couldn't serialize proto response. %v", err)
		return
	}

	// Send response
	err = msg.Respond(responseProto)
	if err != nil {
		log.Errorf("Couldn't send response. %v", err)
	}
}
