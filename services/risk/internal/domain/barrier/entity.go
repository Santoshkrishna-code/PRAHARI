package barrier

import (
	"errors"
)

// BarrierType (Preventive / Mitigative) for bow-tie analysis.
type BarrierType string

const (
	TypePreventive BarrierType = "PREVENTIVE"
	TypeMitigative BarrierType = "MITIGATIVE"
)

// Barrier models bow-tie analysis check barrier blocks.
type Barrier struct {
	ID          string      `json:"id" db:"id"`
	RiskID      string      `json:"risk_id" db:"risk_id"`
	BarrierType BarrierType `json:"barrier_type" db:"barrier_type"`
	Description string      `json:"description" db:"description"`
	IsAssured   bool        `json:"is_assured" db:"is_assured"`
}

// Validate checks domain invariants.
func (b *Barrier) Validate() error {
	if b.RiskID == "" {
		return errors.New("risk ID reference is required")
	}
	if b.Description == "" {
		return errors.New("barrier description cannot be empty")
	}
	return nil
}
