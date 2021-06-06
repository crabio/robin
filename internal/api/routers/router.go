package apirouters

import (
	// External
	log "github.com/sirupsen/logrus"
	// Internal
	apiprocessors "github.com/iakrevetkho/robin/internal/api/processors"
	"github.com/iakrevetkho/robin/internal/api/proto"
)

func RouteMsg(msg *proto.Msg) {
	log.Debugf("Route message UUID:%s ts:", msg.Uuid, msg.Ts)

	switch msg.GetPayload().(type) {
	case *proto.Msg_AuthUserRequest:
		apiprocessors.AuthUserRequest(msg.GetAuthUserRequest())
	default:
		log.Errorf("Unknown message type: %+v", msg)
	}
}
