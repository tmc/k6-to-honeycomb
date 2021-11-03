package k6tohoneycomb

import "time"

// K6DataPoint represents a datapoint streamed from k6.
type K6DataPoint struct {
	Type   string `json:"type,omitempty"`
	Metric string `json:"metric,omitempty"`
	Data   struct {
		Tags  map[string]string `json:"tags,omitempty"`
		Time  time.Time         `json:"time,omitempty"`
		Value float64           `json:"value,omitempty"`
	} `json:"data,omitempty"`
}
