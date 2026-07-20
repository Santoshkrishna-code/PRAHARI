package document

import (
	"errors"
)

// Document tracks contracts certifications references.
type Document struct {
	ID           string `json:"id" db:"id"`
	ContractorID string `json:"contractor_id" db:"contractor_id"`
	Title        string `json:"title" db:"title"`
	StoragePath  string `json:"storage_path" db:"storage_path"`
}

// Validate checks domain invariants.
func (d *Document) Validate() error {
	if d.ContractorID == "" {
		return errors.New("contractor ID reference is required")
	}
	if d.Title == "" {
		return errors.New("document title is required")
	}
	return nil
}
