package model

import "time"

type StatusApiModel struct {
	Message   string    `json:"message,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}
