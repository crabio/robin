package apiprocessors

import (
	// External
	log "github.com/sirupsen/logrus"
	// Internal
	"github.com/iakrevetkho/robin/internal/api/proto"
)

func AuthUserRequest(msg *proto.AuthUserRequest) {
	log.Debugf("Process AuthUserRequest: %+v", msg)

}
