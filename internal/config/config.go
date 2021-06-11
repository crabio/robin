package config

type Config struct {
	NATS struct {
		Hostname string `required:"true" default:"localhost"`
		Request  struct {
			Subject string `default:"auth-request"`
			Queue   string `default:"robin"`
		}
	}
	Auth struct {
		Google struct {
			SecretFileName string `default:"secrets/google.json"`
			RedirectURL    string `default:"http://localhost:9000/auth"`
		}
	}
}
