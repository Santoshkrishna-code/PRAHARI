package snapshot

import "time"

// Image holds frames captured during violation alerts.
type Image struct {
	ID          string    `json:"id"`
	CameraID    string    `json:"camera_id"`
	TriggerID   string    `json:"trigger_id"`
	CapturedAt  time.Time `json:"captured_at"`
	StoragePath string    `json:"storage_path"`
}
