package model

import "time"

type StatusApiModel struct {
	Message   string    `json:"message,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type VersionApiModel struct {
	Version   string `json:"version,omitempty"`
	GitCommit string `json:"gitCommit,omitempty"`
	GitAuthor string `json:"gitAuthor,omitempty"`
	GoVersion string `json:"goVersion,omitempty"`
}

type HealthApiModel struct {
	Status     Status                       `json:"status,omitempty"`
	Components map[string]ComponentApiModel `json:"components,omitempty"`
}

type ComponentApiModel struct {
	Status string `json:"status,omitempty"`
}

type Status struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}
