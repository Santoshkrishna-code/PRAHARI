package comment

import (
	"errors"
	"time"
)

// Comment represents a threaded discussion entry on an incident.
// Comments support nesting via ParentCommentID to allow reply chains.
type Comment struct {
	ID              string    `json:"id" db:"id"`
	IncidentID      string    `json:"incident_id" db:"incident_id"`
	AuthorID        string    `json:"author_id" db:"author_id"`
	ParentCommentID string    `json:"parent_comment_id,omitempty" db:"parent_comment_id"`
	Body            string    `json:"body" db:"body"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	IsDeleted       bool      `json:"is_deleted" db:"is_deleted"`
}

// Validate enforces domain invariants on the comment aggregate.
func (c *Comment) Validate() error {
	if c.IncidentID == "" {
		return errors.New("incident ID is required for comment")
	}
	if c.AuthorID == "" {
		return errors.New("author ID is required for comment")
	}
	if c.Body == "" {
		return errors.New("comment body cannot be empty")
	}
	if len(c.Body) > 10000 {
		return errors.New("comment body must not exceed 10000 characters")
	}
	return nil
}

// IsReply returns true if this comment is a reply to another comment.
func (c *Comment) IsReply() bool {
	return c.ParentCommentID != ""
}

// Edit updates the comment body and records the modification timestamp.
func (c *Comment) Edit(newBody string) {
	c.Body = newBody
	c.UpdatedAt = time.Now()
}
