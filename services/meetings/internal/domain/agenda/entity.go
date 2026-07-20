package agenda

import "time"

// Item represents a single agenda item within a meeting.
type Item struct {
	ID          string    `json:"id"`
	MeetingID   string    `json:"meeting_id"`
	SeqOrder    int       `json:"seq_order"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PresenterID string    `json:"presenter_id,omitempty"`
	DurationMin int       `json:"duration_min"`
	CreatedAt   time.Time `json:"created_at"`
}
