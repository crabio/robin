package config

type Config struct {
	NATS struct {
		Hostname string `required:"true" default:"localhost"`
		Request  struct {
			Subject string `default:"robin"`
			Queue   string `default:"auth-request"`
		}
	}
}
