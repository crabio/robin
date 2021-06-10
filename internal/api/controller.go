package api

import (
	// Internal
	connectors "github.com/iakrevetkho/robin/internal/api/connectors"
	resources "github.com/iakrevetkho/robin/internal/api/resources"
	config "github.com/iakrevetkho/robin/internal/config"
)

// API entities controller
type APIController struct {
	natsConnPtr *connectors.NatsConnector
}

// Init all connectors to systems
func Init(controllerData resources.ControllerData, config config.Config) (cntrl APIController, err error) {
	cntrl.natsConnPtr, err = connectors.NatsInit(controllerData, config)
	if err != nil {
		return
	}

	return
}

func (cntrl *APIController) Close() (err error) {
	err = cntrl.natsConnPtr.Close()
	if err != nil {
		return
	}

	return
}
