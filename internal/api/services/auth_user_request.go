package apiservices

import (
	// External
	log "github.com/sirupsen/logrus"
	// Internal
	auth_google "github.com/iakrevetkho/robin/internal/auth/google"
	resources "github.com/iakrevetkho/robin/internal/resources"
)

func AuthUserRequest(googleAuthProvider *auth_google.Provider, msg *resources.AuthUserRequest) (response *resources.AuthUserResponse, err error) {
	log.Debugf("Process AuthUserRequest: %+v", msg)

	switch msg.Provider {
	case resources.AuthProviderEnum_google:
		// googleAuthProvider.AuthURL
	}

	response = &resources.AuthUserResponse{
		Success: true,
	}
	return
}
