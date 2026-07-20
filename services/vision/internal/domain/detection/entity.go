package detection

import "time"

// Detection represents a bounding box detection event.
type Detection struct {
	ID         string    `json:"id"`
	JobID      string    `json:"job_id"`
	Label      string    `json:"label"` // E.g., no_helmet, smoke, spill
	Confidence float64   `json:"confidence"`
	BBox       string    `json:"bbox"` // JSON bounds E.g., [x, y, w, h]
	Timestamp  time.Time `json:"timestamp"`
}
