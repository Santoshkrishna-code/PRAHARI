package search

import "time"

// Criteria defines parameters to filter historical coordinates and bounding boxes.
type Criteria struct {
	CameraID  string     `json:"camera_id,omitempty"`
	Label     string     `json:"label,omitempty"`
	StartTime *time.Time `json:"start_time,omitempty"`
	EndTime   *time.Time `json:"end_time,omitempty"`
	Limit     int        `json:"limit,omitempty"`
	Offset    int        `json:"offset,omitempty"`
}
