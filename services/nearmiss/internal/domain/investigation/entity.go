package investigation

import (
	"errors"
	"time"
)

// Investigation parameters models.
type Investigation struct {
	ID                 string    `json:"id" db:"id"`
	NearMissID         string    `json:"near_miss_id" db:"near_miss_id"`
	LeadInvestigatorID string    `json:"lead_investigator_id" db:"lead_investigator_id"`
	InvestigationDate  time.Time `json:"investigation_date" db:"investigation_date"`
	Findings           TEXT      `json:"findings" db:"findings"`
	Methodology        string    `json:"methodology" db:"methodology"` // 5 Whys, Fishbone, etc
}

// Validate checks domain invariants.
func (i *Investigation) Validate() error {
	if i.NearMissID == "" {
		return errors.New("near miss ID reference is required")
	}
	if i.LeadInvestigatorID == "" {
		return errors.New("lead investigator ID reference is required")
	}
	return nil
}
