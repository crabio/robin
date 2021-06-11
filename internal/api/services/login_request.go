package apiservices

import (
	// External
	log "github.com/sirupsen/logrus"
	// Internal
	apiresources "github.com/iakrevetkho/robin/internal/api/resources"
	resources "github.com/iakrevetkho/robin/internal/resources"
)

func LoginRequest(controllerData apiresources.ControllerData, msg *resources.LoginRequest) (response *resources.LoginResponse, err error) {
	log.Debugf("Process login request for provider '%s'", msg.GetProvider())

	switch msg.Provider {
	case resources.AuthProviderEnum_google:
		response = &resources.LoginResponse{
			Url: controllerData.GoogleAuthProvider.AuthURL,
		}
	default:
		log.Errorf("Unknown auth provider for login request: '%s'", msg.Provider)
	}

	return
}
