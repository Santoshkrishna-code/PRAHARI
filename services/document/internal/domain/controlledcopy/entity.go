package controlledcopy

import "time"

// Copy tracks watermarked physical or electronic controlled copy issuances.
type Copy struct {
	ID           string    `json:"id"`
	DocumentID   string    `json:"document_id"`
	VersionID    string    `json:"version_id"`
	CopyNumber   string    `json:"copy_number"`
	IssuedTo     string    `json:"issued_to"`
	Location     string    `json:"location"`
	Status       string    `json:"status"` // ACTIVE, RECALLED, DESTROYED
	IssuedAt     time.Time `json:"issued_at"`
}
