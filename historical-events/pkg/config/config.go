package config

type Config struct {
	Server Server
}

type Server struct {
	Port           string
	TrustedProxies []string
}

func NewRandomConfig() Config {
	return Config{
		Server: Server{
			Port:           "8080",
			TrustedProxies: []string{"localhost"},
		},
	}
}
