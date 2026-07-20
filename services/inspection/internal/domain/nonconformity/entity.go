package nonconformity

import (
	"errors"
	"time"
)

// NonConformity represents a formal Non-Conformity Report (NCR) generated from walkthrough findings.
type NonConformity struct {
	ID          string    `json:"id" db:"id"`
	InspectionID string   `json:"inspection_id" db:"inspection_id"`
	FindingID   string    `json:"finding_id" db:"finding_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	IssuerID    string    `json:"issuer_id" db:"issuer_id"`
	ReceiverID  string    `json:"receiver_id" db:"receiver_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks domain invariants.
func (nc *NonConformity) Validate() error {
	if nc.InspectionID == "" {
		return errors.New("inspection ID is required for non-conformity")
	}
	if nc.FindingID == "" {
		return errors.New("finding ID is required for non-conformity")
	}
	if nc.Title == "" {
		return errors.New("non-conformity title is required")
	}
	if nc.IssuerID == "" {
		return errors.New("issuer identifier is required")
	}
	return nil
}
