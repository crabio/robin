package apirouters

import (
	// External
	"fmt"

	log "github.com/sirupsen/logrus"
	// Internal
	apiservices "github.com/iakrevetkho/robin/internal/api/services"
	resources "github.com/iakrevetkho/robin/internal/resources"
)

func RouteMsg(request *resources.Msg) (response *resources.Msg, err error) {
	log.Debugf("Route message UUID:%s", request.Uuid.Value)

	switch request.GetPayload().(type) {
	case *resources.Msg_AuthUserRequest:
		responsePayload, err := apiservices.AuthUserRequest(request.GetAuthUserRequest())
		response := &resources.Msg{
			Uuid: request.Uuid,
			Ts:   request.Ts,
			Payload: &resources.Msg_AuthUserResponse{
				AuthUserResponse: responsePayload,
			},
		}
		return response, err
	default:
		err = fmt.Errorf("Unknown message type: %+v", request)
	}
	return
}
