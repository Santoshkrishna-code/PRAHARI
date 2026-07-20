package playback

import "time"

// Session represents a historical timeline replay session.
type Session struct {
	ID        string    `json:"id"`
	TwinID    string    `json:"twin_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Speed     float64   `json:"speed"` // 1.0x, 2.0x, 0.5x
	Status    string    `json:"status"` // PLAYING, PAUSED, STOPPED
	CreatedAt time.Time `json:"created_at"`
}
