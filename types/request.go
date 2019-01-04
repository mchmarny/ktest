package types

import (
	"time"
)

// SimpleRequest represents simple request
type SimpleRequest struct {
	Meta    *RequestMetadata       `json:"meta"`
	Headers map[string]interface{} `json:"head"`
	EnvVars map[string]interface{} `json:"envs"`
}

// RequestMetadata represents metadata of the request
type RequestMetadata struct {
	ID     string    `json:"id"`
	Ts     time.Time `json:"ts"`
	URI    string    `json:"uri"`
	Host   string    `json:"host"`
	Method string    `json:"method"`
}
