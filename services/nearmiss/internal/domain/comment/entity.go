package comment

import (
	"errors"
	"time"
)

// Comment tracks discussions notes.
type Comment struct {
	ID              string    `json:"id" db:"id"`
	NearMissID      string    `json:"near_miss_id" db:"near_miss_id"`
	AuthorID        string    `json:"author_id" db:"author_id"`
	ParentCommentID string    `json:"parent_comment_id,omitempty" db:"parent_comment_id"`
	Body            string    `json:"body" db:"body"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	IsDeleted       bool      `json:"is_deleted" db:"is_deleted"`
}

// Validate checks invariants.
func (c *Comment) Validate() error {
	if c.NearMissID == "" {
		return errors.New("near miss ID reference is required")
	}
	if c.AuthorID == "" {
		return errors.New("author ID reference is required")
	}
	if c.Body == "" {
		return errors.New("comment body cannot be empty")
	}
	return nil
}
