package connectors

import (
	// External

	// Internal
	"github.com/iakrevetkho/robin/internal/config"
	nats_connector "github.com/iakrevetkho/robin/internal/connectors/nats"
)

// Connector to external systems for processing requests
type Connector struct {
	natsConnPtr *nats_connector.Connector
}

// Init all connectors to systems
func Init(config config.Config) (conn Connector, err error) {
	natsConnector, err := nats_connector.Init(config)
	if err != nil {
		return
	}
	conn.natsConnPtr = &natsConnector

	return
}

func (conn *Connector) Close() {
	conn.natsConnPtr.Close()
}
