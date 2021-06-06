package config

type Config struct {
	NATS struct {
		Hostname string `required:"true" default:"localhost"`
	}
}
