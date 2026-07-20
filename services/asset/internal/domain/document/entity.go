package document

import (
	"errors"
)

// Document tracks engineering datasheets, drawings, and manuals.
type Document struct {
	ID          string `json:"id" db:"id"`
	AssetID     string `json:"asset_id" db:"asset_id"`
	Title       string `json:"title" db:"title"`
	DocType     string `json:"doc_type" db:"doc_type"` // Blueprint, Manual, Datasheet
	StoragePath string `json:"storage_path" db:"storage_path"`
}

// Validate checks domain invariants.
func (d *Document) Validate() error {
	if d.AssetID == "" {
		return errors.New("asset ID is required for documents")
	}
	if d.Title == "" {
		return errors.New("document title is required")
	}
	return nil
}
