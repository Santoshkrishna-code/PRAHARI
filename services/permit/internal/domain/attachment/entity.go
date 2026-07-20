package attachment

import (
	"errors"
	"fmt"
	"time"
)

// MaxFileSizeBytes defines the maximum allowed attachment size (50 MB).
const MaxFileSizeBytes = 50 * 1024 * 1024

// AllowedContentTypes defines the set of permitted MIME types.
var AllowedContentTypes = map[string]bool{
	"image/jpeg":      true,
	"image/png":       true,
	"image/gif":       true,
	"image/webp":      true,
	"application/pdf": true,
	"text/plain":      true,
}

// Attachment represents a file uploaded to substantiate a permit (e.g. photos, diagrams).
type Attachment struct {
	ID          string    `json:"id" db:"id"`
	PermitID    string    `json:"permit_id" db:"permit_id"`
	FileName    string    `json:"file_name" db:"file_name"`
	FileSize    int64     `json:"file_size" db:"file_size"`
	ContentType string    `json:"content_type" db:"content_type"`
	StoragePath string    `json:"storage_path" db:"storage_path"`
	UploadedBy  string    `json:"uploaded_by" db:"uploaded_by"`
	UploadedAt  time.Time `json:"uploaded_at" db:"uploaded_at"`
}

// Validate checks domain invariants for Attachment.
func (a *Attachment) Validate() error {
	if a.PermitID == "" {
		return errors.New("permit ID is required for attachment")
	}
	if a.FileName == "" {
		return errors.New("file name is required")
	}
	if a.FileSize <= 0 {
		return errors.New("file size must be greater than zero")
	}
	if a.FileSize > MaxFileSizeBytes {
		return fmt.Errorf("file size exceeds maximum allowed size of %d bytes", MaxFileSizeBytes)
	}
	if !AllowedContentTypes[a.ContentType] {
		return fmt.Errorf("content type %s is not allowed", a.ContentType)
	}
	return nil
}
