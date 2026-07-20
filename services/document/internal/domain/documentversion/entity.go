package documentversion

import "time"

// Version represents a specific revision of a controlled document.
type Version struct {
	ID            string    `json:"id"`
	DocumentID    string    `json:"document_id"`
	VersionNumber string    `json:"version_number"` // E.g., 1.0, 1.1, 2.0
	FileURL       string    `json:"file_url"`
	FileHash      string    `json:"file_hash"` // SHA-256 integrity hash
	ChangeSummary string    `json:"change_summary"`
	CreatedBy     string    `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
}
