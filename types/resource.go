package types

import (
	"time"
)

// SimpleResource represents simple resource
type SimpleResource struct {
	Meta *RequestMetadata `json:"meta,omitempty"`
	Node *SimpleNodeInfo  `json:"node,omitempty"`
	Pod  *SimplePodInfo   `json:"pod,omitempty"`
}

// SimpleNodeInfo represents host info data
type SimpleNodeInfo struct {
	ID        string              `json:"hostId,omitempty"`
	BootTime  time.Time           `json:"bootTs,omitempty"`
	OS        string              `json:"os,omitempty"`
	Resources *SimpleResourceInfo `json:"resources,omitempty"`
}

// SimplePodInfo represents pod info
type SimplePodInfo struct {
	Hostname string              `json:"hostname,omitempty"`
	Limits   *SimpleResourceInfo `json:"limits,omitempty"`
}

// SimpleResourceInfo represents node cpu
type SimpleResourceInfo struct {
	Memory *SimpleIntMeasurement `json:"memory,omitempty"`
	CPU    *SimpleIntMeasurement `json:"cpu,omitempty"`
}

//SimpleIntMeasurement represents int measurement
type SimpleIntMeasurement struct {
	Value   float64 `json:"value,omitempty"`
	Context string  `json:"context,omitempty"`
}
