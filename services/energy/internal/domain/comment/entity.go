package comment

import (
	"errors"
	"time"
)

// Comment thread log entries.
type Comment struct {
	ID         string    `json:"id" db:"id"`
	TargetType string    `json:"target_type" db:"target_type"` // "METER", "PROGRAM"
	TargetID   string    `json:"target_id" db:"target_id"`
	AuthorID   string    `json:"author_id" db:"author_id"`
	Body       string    `json:"body" db:"body"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// Validate checks body structure.
func (c *Comment) Validate() error {
	if c.TargetID == "" || c.TargetType == "" {
		return errors.New("target reference is required")
	}
	if c.Body == "" {
		return errors.New("comment body is required")
	}
	return nil
}
