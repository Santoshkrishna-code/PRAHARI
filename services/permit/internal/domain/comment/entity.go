package comment

import (
	"errors"
	"time"
)

// Comment represents a threaded discussion entry on a permit.
type Comment struct {
	ID              string    `json:"id" db:"id"`
	PermitID        string    `json:"permit_id" db:"permit_id"`
	AuthorID        string    `json:"author_id" db:"author_id"`
	ParentCommentID string    `json:"parent_comment_id,omitempty" db:"parent_comment_id"`
	Body            string    `json:"body" db:"body"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	IsDeleted       bool      `json:"is_deleted" db:"is_deleted"`
}

// Validate checks domain invariants for Comment.
func (c *Comment) Validate() error {
	if c.PermitID == "" {
		return errors.New("permit ID is required for comment")
	}
	if c.AuthorID == "" {
		return errors.New("author ID is required")
	}
	if c.Body == "" {
		return errors.New("comment body cannot be empty")
	}
	if len(c.Body) > 5000 {
		return errors.New("comment body must not exceed 5000 characters")
	}
	return nil
}

// IsReply checks if this is a sub-thread reply.
func (c *Comment) IsReply() bool {
	return c.ParentCommentID != ""
}
