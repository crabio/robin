package connectorsnats

import (
	// External

	"time"

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"

	// Internal
	resources "github.com/iakrevetkho/robin/internal/resources"
)

func sendErrorResponse(msg *nats.Msg, requestUUID *resources.UUID, err error) {
	// Create message
	response := resources.Msg{
		Uuid: requestUUID,
		Ts: &resources.Timestamp{
			Value: uint64(time.Now().Unix()),
		},
		Payload: &resources.Msg_Error{
			Error: &resources.Error{
				Reason: err.Error(),
			},
		},
	}

	// Serialize response
	responseProto, err := proto.Marshal(&response)
	if err != nil {
		log.Fatalf("Couldn't serialize proto response: %v", err)
	}

	// Send response
	err = msg.Respond(responseProto)
	if err != nil {
		log.Errorf("Couldn't send response. %v", err)
	}
}
