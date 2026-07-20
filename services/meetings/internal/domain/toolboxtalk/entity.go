package toolboxtalk

import "time"

// Talk represents a toolbox talk session — a short, focused safety discussion typically
// conducted at the start of a shift or work period covering a specific safety topic.
type Talk struct {
	ID          string    `json:"id"`
	MeetingID   string    `json:"meeting_id"`
	TopicTitle  string    `json:"topic_title"`
	TopicBody   string    `json:"topic_body"`
	Category    string    `json:"category"` // E.g., HAZMAT, ELECTRICAL, FALL_PROTECTION, CONFINED_SPACE
	Mandatory   bool      `json:"mandatory"`
	FrequencyDays int    `json:"frequency_days"` // Recurrence period in days
	CreatedAt   time.Time `json:"created_at"`
}
