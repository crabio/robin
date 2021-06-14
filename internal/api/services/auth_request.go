package apiservices

import (
	// External
	"fmt"
	"net/url"

	log "github.com/sirupsen/logrus"

	// Internal
	apihelpers "github.com/iakrevetkho/robin/internal/api/helpers"
	"github.com/iakrevetkho/robin/internal/api/resources"
	apiresources "github.com/iakrevetkho/robin/internal/api/resources"
	proto_resources "github.com/iakrevetkho/robin/internal/proto_resources"
)

func AuthRequest(controllerData apiresources.ControllerData, msg *proto_resources.AuthRequest) (response *proto_resources.AuthResponse, err error) {
	log.WithFields(log.Fields{"provider": msg.GetProvider(), "provider_url_response": msg.GetAuthProviderUrlResponse()}).Debug("Process auth request for provider")

	// Parse response URL from the auth provider
	authProviderURL, err := url.Parse(msg.GetAuthProviderUrlResponse())
	if err != nil {
		return nil, err
	}
	log.WithField("authProvider query", authProviderURL.Query()).Debug("Parsed auth provider URL")

	var authProviderResponse = resources.AuthProviderResponse{}
	err = apihelpers.DecodeQueryStringParams(authProviderURL, &authProviderResponse)
	if err != nil {
		return nil, err
	}
	log.WithField("authProviderResponse", authProviderResponse).Debug("Parsed auth provider response params")

	switch msg.Provider {

	case proto_resources.AuthProviderEnum_google:
		userProfile, err := controllerData.GoogleAuthProvider.ProcessAuthRedirect(authProviderResponse.AuthCode)

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
