package apirouters

import (
	// External
	log "github.com/sirupsen/logrus"
	// Internal
	apiservices "github.com/iakrevetkho/robin/internal/api/services"
	resources "github.com/iakrevetkho/robin/internal/resources"
)

func RouteMsg(msg *resources.Msg) {
	log.Debugf("Route message UUID:%s", msg.Uuid.Value)

	switch msg.GetPayload().(type) {
	case *resources.Msg_AuthUserRequest:
		apiservices.AuthUserRequest(msg.GetAuthUserRequest())
	default:
		log.Errorf("Unknown message type: %+v", msg)
	}
}
