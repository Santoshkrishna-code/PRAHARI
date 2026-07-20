package disclosure

import (
	"errors"
	"time"
)

// Disclosure records GRI/SASB compliance reports disclosures.
type Disclosure struct {
	ID             string    `json:"id" db:"id"`
	FrameworkID    string    `json:"framework_id" db:"framework_id"`
	ReferenceCode  string    `json:"reference_code" db:"reference_code"` // e.g. "GRI-302-1"
	DisclosureText string    `json:"disclosure_text" db:"disclosure_text"`
	Status         string    `json:"status" db:"status"` // "DRAFT", "VALIDATED", "PUBLISHED"
	ApprovedByID   string    `json:"approved_by_id" db:"approved_by_id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks fields.
func (d *Disclosure) Validate() error {
	if d.FrameworkID == "" {
		return errors.New("parent framework reference ID is required")
	}
	if d.ReferenceCode == "" {
		return errors.New("disclosure standard reference code is required")
	}
	return nil
}
