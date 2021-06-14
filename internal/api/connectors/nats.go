package apiconnectors

import (
	// External
	"fmt"

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"

	// Internal
	controllers "github.com/iakrevetkho/robin/internal/api/controllers"
	apiresources "github.com/iakrevetkho/robin/internal/api/resources"
	"github.com/iakrevetkho/robin/internal/config"
)

// Connector for processing requests from NATS broker
type NatsConnector struct {
	ConnPtr *nats.Conn
	SubPtr  *nats.Subscription
}

// Create ant init instance of NATS connector
func NatsInit(controllerData apiresources.ControllerData, config config.Config) (conn *NatsConnector, err error) {
	connectionURL := fmt.Sprintf("%s:%d", config.NATS.Hostname, config.NATS.Port)
	log.WithField("url", connectionURL).Debug("Try to connect to NATS")

	// Connect to broker
	connPtr, err := nats.Connect(connectionURL)
	if err != nil {
		return
	}

	// Subscribe on requests queue
	subPtr, err := connPtr.QueueSubscribe(
		config.NATS.Request.Subject,
		config.NATS.Request.Queue,
		func(msg *nats.Msg) {
			err = controllers.NatsMessage(controllerData, msg)
			if err != nil {
				log.Error(err)
				return
			}
		})
	if err != nil {
		// Close connection
		connPtr.Close()
		return
	}

	conn = &NatsConnector{
		ConnPtr: connPtr,
		SubPtr:  subPtr,
	}
	log.WithField("url", connectionURL).Debug("Connected to NATS")

	return
}

func (conn *NatsConnector) Close() (err error) {
	// Close subscription
	err = conn.SubPtr.Unsubscribe()
	// Close connection
	conn.ConnPtr.Close()

	return
}
