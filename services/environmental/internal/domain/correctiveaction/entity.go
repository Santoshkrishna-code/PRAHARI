package correctiveaction

import (
	"errors"
	"time"
)

// CorrectiveAction structures CAPA workflows.
type CorrectiveAction struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	SourceType     string    `json:"source_type" db:"source_type"` // "SPILL", "EMISSION_EXCEEDANCE", "COMPLIANCE_AUDIT"
	SourceID       string    `json:"source_id" db:"source_id"`
	Description    string    `json:"description" db:"description"`
	AssignedTo     string    `json:"assigned_to" db:"assigned_to"`
	TargetDate     time.Time `json:"target_date" db:"target_date"`
	ActualDate     time.Time `json:"actual_date" db:"actual_date"`
	Status         string    `json:"status" db:"status"` // "OPEN", "VERIFICATION", "CLOSED"
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks action variables.
func (c *CorrectiveAction) Validate() error {
	if c.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	if c.Description == "" {
		return errors.New("corrective action description is required")
	}
	return nil
}
