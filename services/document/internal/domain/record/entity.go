package record

import "time"

// Record represents an archived or immutable document record.
type Record struct {
	ID           string    `json:"id"`
	DocumentID   string    `json:"document_id"`
	ArchivedAt   time.Time `json:"archived_at"`
	ArchivedBy   string    `json:"archived_by"`
	StorageClass string    `json:"storage_class"` // GLACIER, DEEP_ARCHIVE
}
