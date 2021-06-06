package apiservices

import (
	// External
	log "github.com/sirupsen/logrus"
	// Internal
	resources "github.com/iakrevetkho/robin/internal/resources"
)

func AuthUserRequest(msg *resources.AuthUserRequest) (response *resources.AuthUserResponse, err error) {
	log.Debugf("Process AuthUserRequest: %+v", msg)

	response = &resources.AuthUserResponse{
		Success: true,
	}
	return
}
