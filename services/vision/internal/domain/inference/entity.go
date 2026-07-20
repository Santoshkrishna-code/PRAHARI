package inference

import "time"

// Job represents a running video stream inference pipeline schedule.
type Job struct {
	ID        string    `json:"id"`
	CameraID  string    `json:"camera_id"`
	ModelID   string    `json:"model_id"`
	Status    string    `json:"status"` // STARTING, RUNNING, STOPPED, FAILED
	FPSRate   float64   `json:"fps_rate"`
	StartedAt time.Time `json:"started_at"`
}
