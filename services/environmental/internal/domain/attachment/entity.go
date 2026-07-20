package attachment

import (
	"errors"
	"time"
)

// Attachment maps documents.
type Attachment struct {
	ID         string    `json:"id" db:"id"`
	TargetType string    `json:"target_type" db:"target_type"` // "PERMIT", "SAMPLING", "SPILL"
	TargetID   string    `json:"target_id" db:"target_id"`
	FileName   string    `json:"file_name" db:"file_name"`
	FileURL    string    `json:"file_url" db:"file_url"`
	UploadedBy string    `json:"uploaded_by" db:"uploaded_by"`
	UploadedAt time.Time `json:"uploaded_at" db:"uploaded_at"`
}

// Validate checks fields.
func (a *Attachment) Validate() error {
	if a.TargetID == "" || a.TargetType == "" {
		return errors.New("target reference is required")
	}
	if a.FileURL == "" {
		return errors.New("file storage URL is required")
	}
	return nil
}
