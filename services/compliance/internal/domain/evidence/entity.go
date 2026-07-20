package evidence

import (
	"errors"
	"time"
)

// Evidence logs compliance documentation upload files reference tags.
type Evidence struct {
	ID           string    `json:"id" db:"id"`
	ObligationID string    `json:"obligation_id" db:"obligation_id"`
	UploadedByID string    `json:"uploaded_by_id" db:"uploaded_by_id"`
	StoragePath  string    `json:"storage_path" db:"storage_path"`
	CollectedAt  time.Time `json:"collected_at" db:"collected_at"`
}

// Validate checks domain invariants.
func (e *Evidence) Validate() error {
	if e.ObligationID == "" {
		return errors.New("obligation ID reference is required")
	}
	if e.StoragePath == "" {
		return errors.New("evidence storage path cannot be empty")
	}
	return nil
}
