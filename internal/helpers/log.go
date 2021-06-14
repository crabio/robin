package helpers

import (
	// External
	"os"

	log "github.com/sirupsen/logrus"
	// Internal
)

func InitLog() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}
