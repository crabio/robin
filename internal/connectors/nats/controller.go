package connectorsnats

import (
	// External
	"fmt"

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"

	// Internal
	apirouters "github.com/iakrevetkho/robin/internal/api/routers"
	resources "github.com/iakrevetkho/robin/internal/resources"
)

func processMsg(msg *nats.Msg) {
	// Check that msg is request
	if msg.Reply == "" {
		err := fmt.Errorf("Await request, but receive msg without reply: %+v", msg)
		log.Error(err)
		sendErrorResponse(msg, &resources.UUID{}, err)
		return
	}

	// Parse message protobuf
	protoMsg := resources.Msg{}
	if err := proto.Unmarshal(msg.Data, &protoMsg); err != nil {
		err := fmt.Errorf("Failed to parse proto msg '%s': %v", msg.Data, err)
		log.Error(err)
		sendErrorResponse(msg, &resources.UUID{}, err)
		return
	}

	// Go to msg router
	response, err := apirouters.RouteMsg(&protoMsg)
	if err != nil {
		err := fmt.Errorf("Couldn't process msg '%+v': %v", protoMsg.GetPayload(), err)
		log.Error(err)
		sendErrorResponse(msg, &resources.UUID{}, err)
		return
	}

	// Serialize response
	responseProto, err := proto.Marshal(response)
	if err != nil {
		log.Fatalf("Couldn't serialize proto response: %v", err)
	}

	// Send response
	err = msg.Respond(responseProto)
	if err != nil {
		log.Errorf("Couldn't send response. %v", err)
	}
}
