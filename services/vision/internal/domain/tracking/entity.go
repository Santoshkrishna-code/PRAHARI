package tracking

import "time"

// Segment tracks movements coordinates mapping over time.
type Segment struct {
	ID        string    `json:"id"`
	ObjectID  string    `json:"object_id"`
	CameraID  string    `json:"camera_id"`
	XVal      float64   `json:"x_val"`
	YVal      float64   `json:"y_val"`
	Timestamp time.Time `json:"timestamp"`
}
