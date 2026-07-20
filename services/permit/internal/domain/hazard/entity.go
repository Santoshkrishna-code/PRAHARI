package hazard

import (
	"errors"
	"time"
)

// Type defines safety hazard categories.
type Type string

const (
	TypeFire          Type = "FIRE"
	TypeExplosion     Type = "EXPLOSION"
	TypeToxic         Type = "TOXIC"
	TypeAsphyxiation  Type = "ASPHYXIATION"
	TypeElectrical    Type = "ELECTRICAL"
	TypeMechanical    Type = "MECHANICAL"
	TypeFall          Type = "FALL"
	TypeRadiation     Type = "RADIATION"
	TypeChemical      Type = "CHEMICAL"
	TypeEnvironmental Type = "ENVIRONMENTAL"
)

// Hazard tracks a potential hazard identified in risk assessment.
type Hazard struct {
	ID             string    `json:"id" db:"id"`
	PermitID       string    `json:"permit_id" db:"permit_id"`
	Type           Type      `json:"type" db:"type"`
	Description    string    `json:"description" db:"description"`
	ControlMeasure string    `json:"control_measure" db:"control_measure"`
	IdentifiedBy   string    `json:"identified_by" db:"identified_by"`
	IdentifiedAt   time.Time `json:"identified_at" db:"identified_at"`
}

// Validate checks domain invariants for Hazard.
func (h *Hazard) Validate() error {
	if h.PermitID == "" {
		return errors.New("permit ID is required for hazard")
	}
	if h.Type == "" {
		return errors.New("hazard type is required")
	}
	if h.Description == "" {
		return errors.New("hazard description is required")
	}
	if h.IdentifiedBy == "" {
		return errors.New("identifier identity is required")
	}
	return nil
}
