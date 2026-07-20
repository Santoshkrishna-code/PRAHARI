package comment

import (
	"errors"
	"time"
)

// Comment tracks audit discussions on inspections.
type Comment struct {
	ID              string    `json:"id" db:"id"`
	InspectionID    string    `json:"inspection_id" db:"inspection_id"`
	AuthorID        string    `json:"author_id" db:"author_id"`
	ParentCommentID string    `json:"parent_comment_id,omitempty" db:"parent_comment_id"`
	Body            string    `json:"body" db:"body"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	IsDeleted       bool      `json:"is_deleted" db:"is_deleted"`
}

// Validate checks domain invariants for Comment.
func (c *Comment) Validate() error {
	if c.InspectionID == "" {
		return errors.New("inspection ID is required for comment")
	}
	if c.AuthorID == "" {
		return errors.New("author ID is required")
	}
	if c.Body == "" {
		return errors.New("comment body cannot be empty")
	}
	return nil
}

// IsReply checks reply chain positions.
func (c *Comment) IsReply() bool {
	return c.ParentCommentID != ""
}
