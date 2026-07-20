package mitigation

import (
	"errors"
	"time"
)

// Mitigation maps plans steps.
type Mitigation struct {
	ID                   string    `json:"id" db:"id"`
	HazardID             string    `json:"hazard_id" db:"hazard_id"`
	Description          string    `json:"description" db:"description"`
	TargetCompletionDate time.Time `json:"target_completion_date" db:"target_completion_date"`
	ResponsiblePartyID   string    `json:"responsible_party_id" db:"responsible_party_id"`
	IsImplemented        bool      `json:"is_implemented" db:"is_implemented"`
}

// Validate checks domain invariants.
func (m *Mitigation) Validate() error {
	if m.HazardID == "" {
		return errors.New("hazard ID is required")
	}
	if m.Description == "" {
		return errors.New("mitigation plan details description is required")
	}
	return nil
}
