package preventiveaction

import (
	"errors"
	"time"
)

// PreventiveAction details CAPA preventive actions.
type PreventiveAction struct {
	ID                 string    `json:"id" db:"id"`
	NearMissID         string    `json:"near_miss_id" db:"near_miss_id"`
	Description        string    `json:"description" db:"description"`
	TargetDate         time.Time `json:"target_date" db:"target_date"`
	ResponsiblePartyID string    `json:"responsible_party_id" db:"responsible_party_id"`
	IsImplemented      bool      `json:"is_implemented" db:"is_implemented"`
}

// Validate checks domain invariants.
func (pa *PreventiveAction) Validate() error {
	if pa.NearMissID == "" {
		return errors.New("near miss ID is required")
	}
	if pa.Description == "" {
		return errors.New("preventive action description is required")
	}
	return nil
}
