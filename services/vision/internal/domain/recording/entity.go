package recording

import "time"

// Session records stream files.
type Session struct {
	ID        string    `json:"id"`
	CameraID  string    `json:"camera_id"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	FileURL   string    `json:"file_url"`
}
