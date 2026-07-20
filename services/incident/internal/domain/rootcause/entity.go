package rootcause

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// Category classifies the root cause origin.
type Category string

const (
	CategoryHuman         Category = "HUMAN"
	CategoryProcess       Category = "PROCESS"
	CategoryEquipment     Category = "EQUIPMENT"
	CategoryEnvironmental Category = "ENVIRONMENTAL"
	CategoryManagement    Category = "MANAGEMENT"
	CategoryDesign        Category = "DESIGN"
)

// ValidCategories enumerates all accepted root cause categories.
var ValidCategories = []Category{
	CategoryHuman,
	CategoryProcess,
	CategoryEquipment,
	CategoryEnvironmental,
	CategoryManagement,
	CategoryDesign,
}

// RootCause represents the identified root cause of an incident, linked to an investigation.
type RootCause struct {
	ID                  string          `json:"id" db:"id"`
	InvestigationID     string          `json:"investigation_id" db:"investigation_id"`
	IncidentID          string          `json:"incident_id" db:"incident_id"`
	Category            Category        `json:"category" db:"category"`
	Description         string          `json:"description" db:"description"`
	ContributingFactors json.RawMessage `json:"contributing_factors" db:"contributing_factors"`
	IdentifiedBy        string          `json:"identified_by" db:"identified_by"`
	IdentifiedAt        time.Time       `json:"identified_at" db:"identified_at"`
}

// Validate enforces domain invariants on the root cause aggregate.
func (rc *RootCause) Validate() error {
	if rc.InvestigationID == "" {
		return errors.New("investigation ID is required for root cause")
	}
	if rc.IncidentID == "" {
		return errors.New("incident ID is required for root cause")
	}
	if !rc.isValidCategory() {
		return fmt.Errorf("invalid root cause category: %s", rc.Category)
	}
	if rc.Description == "" {
		return errors.New("root cause description is required")
	}
	if rc.IdentifiedBy == "" {
		return errors.New("identifier (identified_by) is required")
	}
	return nil
}

// GetContributingFactors deserializes the contributing factors JSON array.
func (rc *RootCause) GetContributingFactors() ([]string, error) {
	if rc.ContributingFactors == nil {
		return nil, nil
	}
	var factors []string
	if err := json.Unmarshal(rc.ContributingFactors, &factors); err != nil {
		return nil, fmt.Errorf("failed to parse contributing factors: %w", err)
	}
	return factors, nil
}

// isValidCategory checks whether the category is among accepted classifications.
func (rc *RootCause) isValidCategory() bool {
	for _, c := range ValidCategories {
		if rc.Category == c {
			return true
		}
	}
	return false
}
