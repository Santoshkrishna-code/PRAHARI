package alert

import "time"

// EventTrigger represents triggered warnings when detections violate safety rules.
type EventTrigger struct {
	ID          string    `json:"id"`
	CameraID    string    `json:"camera_id"`
	Label       string    `json:"label"` // E.g., no_helmet, smoke
	TriggeredAt time.Time `json:"triggered_at"`
	SnapshotURL string    `json:"snapshot_url,omitempty"`
}
