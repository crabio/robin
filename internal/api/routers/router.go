package apirouters

import (
	// External
	log "github.com/sirupsen/logrus"
	// Internal
	"github.com/iakrevetkho/robin/internal/api/proto"
	apiservices "github.com/iakrevetkho/robin/internal/api/services"
)

func RouteMsg(msg *proto.Msg) {
	log.Debugf("Route message UUID:%s", msg.Uuid.Value)

	switch msg.GetPayload().(type) {
	case *proto.Msg_AuthUserRequest:
		apiservices.AuthUserRequest(msg.GetAuthUserRequest())
	default:
		log.Errorf("Unknown message type: %+v", msg)
	}
}
