package main

import (
	// External
	log "github.com/sirupsen/logrus"
	// Internal
	"github.com/iakrevetkho/robin/internal/config"
)

//Execution starts from main function
func main() {
	// Load app configuration
	config, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Couldn't load app config. %v", err)
	}
	log.Infof("Loaded config: %+v", config)
}
