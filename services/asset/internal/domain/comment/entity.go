package comment

import (
	"errors"
	"time"
)

// Comment tracks audit discussions notes on equipment profiles.
type Comment struct {
	ID              string    `json:"id" db:"id"`
	AssetID         string    `json:"asset_id" db:"asset_id"`
	AuthorID        string    `json:"author_id" db:"author_id"`
	ParentCommentID string    `json:"parent_comment_id,omitempty" db:"parent_comment_id"`
	Body            string    `json:"body" db:"body"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	IsDeleted       bool      `json:"is_deleted" db:"is_deleted"`
}

// Validate checks invariants.
func (c *Comment) Validate() error {
	if c.AssetID == "" {
		return errors.New("asset ID is required for comments")
	}
	if c.AuthorID == "" {
		return errors.New("author ID is required")
	}
	if c.Body == "" {
		return errors.New("comment body cannot be empty")
	}
	return nil
}
