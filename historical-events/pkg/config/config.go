package config

type Config struct {
	Server Server
}

type Server struct {
	Host           string
	Port           string
	TrustedProxies []string
}

func NewRandomConfig() Config {
	return Config{
		Server: Server{
			Host:           "localHost",
			Port:           "8080",
			TrustedProxies: []string{},
		},
	}
}
