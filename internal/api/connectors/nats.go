package apiconnectors

import (
	// External

	"github.com/nats-io/nats.go"

	// Internal
	controllers "github.com/iakrevetkho/robin/internal/api/controllers"
	apiresources "github.com/iakrevetkho/robin/internal/api/resources"
	"github.com/iakrevetkho/robin/internal/config"
)

// Connector for processing requests from NATS broker
type NatsConnector struct {
	NatsConnPtr *nats.Conn
	SubPtr      *nats.Subscription
}

// Create ant init instance of NATS connector
func NatsInit(controllerData apiresources.ControllerData, config config.Config) (conn *NatsConnector, err error) {
	// Connect to broker
	conn.NatsConnPtr, err = nats.Connect(config.NATS.Hostname)
	if err != nil {
		return
	}

	// Subscribe on requests queue
	conn.SubPtr, err = conn.NatsConnPtr.QueueSubscribe(
		config.NATS.Request.Subject,
		config.NATS.Request.Queue,
		func(msg *nats.Msg) {
			controllers.NatsMessage(controllerData, msg)
		})
	if err != nil {
		// Close connection
		conn.NatsConnPtr.Close()
		return
	}

	return
}

func (conn *NatsConnector) Close() (err error) {
	// Close subscription
	err = conn.SubPtr.Unsubscribe()
	// Close connection
	conn.NatsConnPtr.Close()

	return
}
