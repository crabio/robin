package apihelpers

import (
	// External
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"

	// Internal
	resources "github.com/iakrevetkho/robin/internal/resources"
)

// Method for sending error response onto message from NATS broker.
// This method also logs error message bases on `format` and `..args`
func NatsErrorResponse(msg *nats.Msg, requestUUID *resources.UUID, format string, args ...interface{}) error {
	// Log error
	err := fmt.Errorf(format, args...)
	log.Error(err)

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
		return fmt.Errorf("Couldn't serialize proto response: %v", err)
	}

	// Send response
	err = msg.Respond(responseProto)
	if err != nil {
		return fmt.Errorf("Couldn't send response. %v", err)
	}
}
