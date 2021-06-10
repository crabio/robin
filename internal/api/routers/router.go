package apirouters

import (
	// External
	"fmt"

	log "github.com/sirupsen/logrus"

	// Internal
	apiresources "github.com/iakrevetkho/robin/internal/api/resources"
	services "github.com/iakrevetkho/robin/internal/api/services"
	resources "github.com/iakrevetkho/robin/internal/resources"
)

func RouteMsg(controllerData apiresources.ControllerData, request *resources.Msg) (response *resources.Msg, err error) {
	log.Debugf("Route message UUID:%s", request.Uuid.Value)

	switch request.GetPayload().(type) {
	case *resources.Msg_LoginRequest:
		responsePayload, err := services.LoginRequest(controllerData, request.GetLoginRequest())
		response := &resources.Msg{
			Uuid: request.Uuid,
			Ts:   request.Ts,
			Payload: &resources.Msg_LoginResponse{
				LoginResponse: responsePayload,
			},
		}
		return response, err
	default:
		err = fmt.Errorf("Unknown message type: %+v", request)
	}
	return
}
