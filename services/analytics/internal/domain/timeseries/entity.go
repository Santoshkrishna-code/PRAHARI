package timeseries

import "time"

// DataPoint represents a timestamped metric value in time series charts.
type DataPoint struct {
	ID        string    `json:"id"`
	MetricKey string    `json:"metric_key"`
	Val       float64   `json:"val"`
	Timestamp time.Time `json:"timestamp"`
}
