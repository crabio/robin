package connectorsnats

import (
	// External

	"github.com/nats-io/nats.go"

	// Internal
	"github.com/iakrevetkho/robin/internal/config"
)

// Connector for processing requests from NATS broker
type Connector struct {
	NatsConnPtr *nats.Conn
	SubPtr      *nats.Subscription
}

// Create ant init instance of NATS connector
func Init(config config.Config) (conn Connector, err error) {
	// Connect to broker
	conn.NatsConnPtr, err = nats.Connect(config.NATS.Hostname)
	if err != nil {
		return
	}

	// Subscribe on requests queue
	conn.SubPtr, err = conn.NatsConnPtr.QueueSubscribe(config.NATS.Request.Subject, config.NATS.Request.Queue, processMsg)
	if err != nil {
		// Close connection
		conn.NatsConnPtr.Close()
		return
	}

	return
}

func (conn *Connector) Close() {
	// Close subscription
	conn.SubPtr.Unsubscribe()
	// Close connection
	conn.NatsConnPtr.Close()
}
