package correctiveaction

import (
	"errors"
	"time"
)

// CorrectiveAction details CAPA items.
type CorrectiveAction struct {
	ID                 string    `json:"id" db:"id"`
	NearMissID         string    `json:"near_miss_id" db:"near_miss_id"`
	Description        string    `json:"description" db:"description"`
	TargetDate         time.Time `json:"target_date" db:"target_date"`
	ResponsiblePartyID string    `json:"responsible_party_id" db:"responsible_party_id"`
	IsImplemented      bool      `json:"is_implemented" db:"is_implemented"`
}

// Validate checks domain invariants.
func (ca *CorrectiveAction) Validate() error {
	if ca.NearMissID == "" {
		return errors.New("near miss ID reference is required")
	}
	if ca.Description == "" {
		return errors.New("corrective action description cannot be empty")
	}
	return nil
}
