package apiservices

import (
	// External
	"fmt"

	log "github.com/sirupsen/logrus"
	// Internal
	apiresources "github.com/iakrevetkho/robin/internal/api/resources"
	proto_resources "github.com/iakrevetkho/robin/internal/proto_resources"
)

func AuthRequest(controllerData apiresources.ControllerData, msg *proto_resources.AuthRequest) (response *proto_resources.AuthResponse, err error) {
	log.Debugf("Process auth request for provider '%s' with code '%s'", msg.GetProvider(), msg.GetAuthCode())

	switch msg.Provider {

	case proto_resources.AuthProviderEnum_google:
		userProfile, err := controllerData.GoogleAuthProvider.ProcessAuthRedirect(msg.AuthCode)

		if err != nil {
			return nil, err
		} else {
			response = &proto_resources.AuthResponse{
				FirstName: userProfile.FirstName,
				LastName:  userProfile.LastName,
				Email:     userProfile.Email,
				Locale:    userProfile.Locale,
			}
			return response, nil
		}

	default:
		return nil, fmt.Errorf("Unknown auth provider for login request: '%s'", msg.Provider)
	}
}
