package types

import (
	"time"
)

// NewPodInfo returns fully hydrated pod
func NewPodInfo() *SimplePodInfo {
	return &SimplePodInfo{
		Limits: &SimpleResourceInfo{
			RAM: &SimpleMeasurement{},
			CPU: &SimpleMeasurement{},
		},
	}
}

// NewNodeInfo returns fully hydrated node
func NewNodeInfo() *SimpleNodeInfo {
	return &SimpleNodeInfo{
		Resources: &SimpleResourceInfo{
			RAM: &SimpleMeasurement{},
			CPU: &SimpleMeasurement{},
		},
	}
}

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
	RAM *SimpleMeasurement `json:"ram,omitempty"`
	CPU *SimpleMeasurement `json:"cpu,omitempty"`
	GPU *SimpleMeasurement `json:"gpu,omitempty"`
}

//SimpleMeasurement represents int measurement
type SimpleMeasurement struct {
	Value   float64 `json:"value,omitempty"`
	Context string  `json:"context,omitempty"`
}
