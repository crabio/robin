package apirouters

import (
	// External
	log "github.com/sirupsen/logrus"
	// Internal
	"github.com/iakrevetkho/robin/internal/api/proto"
)

func RouteMsg(msg *proto.Msg) (err error) {
	log.Debugf("Route message: %+v", msg)

	switch msg.GetPayload() {
	case proto.Msg_AuthUserRequest:

	}
}
