package main

import (
	// External

	log "github.com/sirupsen/logrus"

	// Internal
	api "github.com/iakrevetkho/robin/internal/api"
	apiresources "github.com/iakrevetkho/robin/internal/api/resources"
	auth_google "github.com/iakrevetkho/robin/internal/auth/google"
	"github.com/iakrevetkho/robin/internal/config"
	"github.com/iakrevetkho/robin/internal/helpers"
)

func main() {
	helpers.InitLog()

	// Load app configuration
	config, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Couldn't load app config. %v", err)
	}
	log.SetLevel(config.LogLevel)
	log.WithField("config", config).Infof("Loaded config")

	// Init auth providers
	googleAuthProvider, err := auth_google.New(config)
	if err != nil {
		log.Fatalf("Couldn't init google auth provider. %v", err)
	}

	// Create api controller data
	controllerData := apiresources.ControllerData{
		GoogleAuthProvider: googleAuthProvider,
	}

	// Init API controller
	apiController, err := api.Init(controllerData, config)
	if err != nil {
		log.Fatalf("Couldn't init API controller. %v", err)
	}
	defer apiController.Close()

	// Wait any terminate signal
	signal := helpers.WaitTermSignals()
	log.Infof("Exit. Catched signal '%s'", signal.String())
}
