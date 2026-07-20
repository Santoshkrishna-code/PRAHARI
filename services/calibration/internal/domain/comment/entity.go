package comment

import "time"

// Comment represents calibration technician notes, supervisor approval comments, or out-of-tolerance impact notes.
type Comment struct {
	ID         string    `json:"id"`
	TargetType string    `json:"target_type"`
	TargetID   string    `json:"target_id"`
	AuthorID   string    `json:"author_id"`
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"created_at"`
}
