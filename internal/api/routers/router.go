package apirouters

import (
	// External
	"fmt"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	// Internal
	apiresources "github.com/iakrevetkho/robin/internal/api/resources"
	services "github.com/iakrevetkho/robin/internal/api/services"
	proto_resources "github.com/iakrevetkho/robin/internal/proto_resources"
)

func RouteMsg(controllerData apiresources.ControllerData, request *proto_resources.Msg) (response *proto_resources.Msg, err error) {
	msgUUID, err := uuid.FromBytes(request.Uuid.Value)
	if err != nil {
		return
	}
	log.Debugf("Route message UUID: %v", msgUUID)

	// Create base message
	response = &proto_resources.Msg{
		Uuid: request.Uuid,
		Ts:   request.Ts,
	}

	switch request.GetPayload().(type) {

	case *proto_resources.Msg_LoginRequest:
		responsePayload, err := services.LoginRequest(controllerData, request.GetLoginRequest())

		// TODO Optimize duplicated code
		if err != nil {
			response.Payload = &proto_resources.Msg_Error{
				Error: &proto_resources.Error{
					Reason: err.Error(),
				},
			}
		} else {
			response.Payload = &proto_resources.Msg_LoginResponse{
				LoginResponse: responsePayload,
			}
		}

		return response, err

	case *proto_resources.Msg_AuthRequest:
		responsePayload, err := services.AuthRequest(controllerData, request.GetAuthRequest())

		// TODO Optimize duplicated code
		if err != nil {
			response.Payload = &proto_resources.Msg_Error{
				Error: &proto_resources.Error{
					Reason: err.Error(),
				},
			}
		} else {
			response.Payload = &proto_resources.Msg_AuthResponse{
				AuthResponse: responsePayload,
			}
		}

		return response, err

	default:
		err = fmt.Errorf("Unknown message type: %+v", request)
	}
	return
}
