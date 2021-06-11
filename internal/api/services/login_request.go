package apiservices

import (
	// External
	"fmt"

	log "github.com/sirupsen/logrus"
	// Internal
	apiresources "github.com/iakrevetkho/robin/internal/api/resources"
	proto_resources "github.com/iakrevetkho/robin/internal/proto_resources"
)

func LoginRequest(controllerData apiresources.ControllerData, msg *proto_resources.LoginRequest) (response *proto_resources.LoginResponse, err error) {
	log.Debugf("Process login request for provider '%s'", msg.GetProvider())

	switch msg.Provider {

	case proto_resources.AuthProviderEnum_google:
		response = &proto_resources.LoginResponse{
			Url: controllerData.GoogleAuthProvider.GetAuthURL(),
		}
		return response, nil

	default:
		return nil, fmt.Errorf("Unknown auth provider for login request: '%s'", msg.Provider)
	}
}
