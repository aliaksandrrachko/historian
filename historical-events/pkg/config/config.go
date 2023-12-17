package config

import "time"

type Config struct {
	Server     Server
	Management Management
}

type Server struct {
	Port           string
	TrustedProxies []string
}

type Management struct {
	Health Health
}

type Health struct {
	GCMaxPauseCheck    time.Duration
	GoRoutineThreshold int
}

func NewRandomConfig() Config {
	return Config{
		Server: Server{
			Port:           "8080",
			TrustedProxies: []string{"127.0.0.1"},
		},
		Management: Management{
			Health: Health{
				GCMaxPauseCheck:    time.Minute,
				GoRoutineThreshold: 100,
			},
		},
	}
}
