package apiconnectors

import (
	// External

	"github.com/nats-io/nats.go"

	// Internal
	apicontrollers "github.com/iakrevetkho/robin/internal/api/controllers"
	"github.com/iakrevetkho/robin/internal/config"
)

// Connector for processing requests from NATS broker
type NatsConnector struct {
	NatsConn *nats.Conn
	Sub      *nats.Subscription
}

// Create ant init instance of NATS connector
func Init(config config.Config) (conn NatsConnector, err error) {
	// Connect to broker
	conn.NatsConn, err = nats.Connect(config.NATS.Hostname)
	if err != nil {
		return
	}

	// Subscribe on requests queue
	conn.Sub, err = conn.NatsConn.QueueSubscribe(config.NATS.Request.Subject, config.NATS.Request.Queue, apicontrollers.ProcessNatsMsg)
	if err != nil {
		// Close connection
		conn.NatsConn.Close()
		return
	}

	return
}

func (conn *NatsConnector) Close() {
	// Close subscription
	conn.Sub.Unsubscribe()
	// Close connection
	conn.NatsConn.Close()
}
