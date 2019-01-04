package types

import (
	"time"
)

// SimpleNode represents simple request
type SimpleNode struct {
	Meta *RequestMetadata `json:"meta"`
	Info *SimpleNodeInfo  `json:"info"`
}

// SimpleNodeInfo represents host info data
type SimpleNodeInfo struct {
	ID       string        `json:"hostId"`
	BootTime time.Time     `json:"bootTs"`
	OS       string        `json:"os"`
	Hostname string        `json:"hostname"`
	Memory   *SimpleMemory `json:"mem"`
}

// SimpleMemory represents node memory
type SimpleMemory struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
}
