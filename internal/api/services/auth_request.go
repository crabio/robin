package apiservices

import (
	// External
	log "github.com/sirupsen/logrus"
	// Internal
	apiresources "github.com/iakrevetkho/robin/internal/api/resources"
	resources "github.com/iakrevetkho/robin/internal/resources"
)

func AuthRequest(controllerData apiresources.ControllerData, msg *resources.AuthRequest) (response *resources.AuthResponse, err error) {
	log.Debugf("Process login request: %+v", msg)

	switch msg.Provider {
	case resources.AuthProviderEnum_google:
		err = controllerData.GoogleAuthProvider.ProcessAuthRedirect(msg.AuthCode)

		if err != nil {
			response = &resources.AuthResponse{
				Success: false,
			}
		} else {
			response = &resources.AuthResponse{
				Success: true,
				Error: &resources.Error{
					Reason: err.Error(),
				},
			}
		}
	default:
		log.Fatalf("Unknown auth provider for login request: '%s'", msg.Provider)
	}

	return
}
