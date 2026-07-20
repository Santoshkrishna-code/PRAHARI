package competencymatrix

import (
	"errors"
)

// CompetencyMatrix maps required roles competencies.
type CompetencyMatrix struct {
	ID           string `json:"id" db:"id"`
	RoleID       string `json:"role_id" db:"role_id"`
	CompetencyID string `json:"competency_id" db:"competency_id"`
}

// Validate checks domain invariants.
func (cm *CompetencyMatrix) Validate() error {
	if cm.RoleID == "" {
		return errors.New("role ID reference is required")
	}
	if cm.CompetencyID == "" {
		return errors.New("competency ID reference is required")
	}
	return nil
}
