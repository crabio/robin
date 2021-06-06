package apiservices

import (
	// External
	log "github.com/sirupsen/logrus"
	// Internal
	resources "github.com/iakrevetkho/robin/internal/resources"
)

func AuthUserRequest(msg *resources.AuthUserRequest) {
	log.Debugf("Process AuthUserRequest: %+v", msg)

}
