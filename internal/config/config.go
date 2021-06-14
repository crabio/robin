package config

import (
	// External
	log "github.com/sirupsen/logrus"
	// Internal
)

// TODO Add logging level
type Config struct {
	LogLevel log.Level `default:"info" env:"LOG_LEVEL"`
	NATS     struct {
		Hostname string `default:"localhost" env:"NATS_HOSTNAME"`
		Port     uint64 `default:"4222" env:"NATS_PORT"`
		Request  struct {
			Subject string `default:"auth-request"`
			Queue   string `default:"robin"`
		}
	}
	SecretsFolderPath string `default:"secrets" env:"SECRETS_FOLDER_PATH"`
	Auth              struct {
		Google struct {
			SecretFileName string `default:"google.json"`
			RedirectURL    string `default:"http://localhost:9000/auth" env:"GOOGLE_REDIRECT_URL"`
		}
	}
}
