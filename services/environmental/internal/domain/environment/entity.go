package environment

import (
	"errors"
	"time"
)

// EnvironmentalAspect defines an element of an organization's activities that interacts with the environment.
type EnvironmentalAspect struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	DepartmentID   string    `json:"department_id" db:"department_id"`
	Name           string    `json:"name" db:"name"` // e.g. "Chemical Storage", "Diesel Generator Stack"
	Description    string    `json:"description" db:"description"`
	AspectCategory string    `json:"aspect_category" db:"aspect_category"` // "AIR_EMISSION", "WASTEWATER", "SOLID_WASTE", "NOISE"
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// EnvironmentalImpact registers the change to the environment resulting from an aspect.
type EnvironmentalImpact struct {
	ID          string `json:"id" db:"id"`
	AspectID    string `json:"aspect_id" db:"aspect_id"`
	Description string `json:"description" db:"description"`
	Severity    int    `json:"severity" db:"severity"`     // 1 to 5 scale
	Probability int    `json:"probability" db:"probability"` // 1 to 5 scale
	RiskScore   int    `json:"risk_score" db:"risk_score"`   // Severity * Probability
}

// Validate checks aspects attributes.
func (a *EnvironmentalAspect) Validate() error {
	if a.PlantID == "" || a.DepartmentID == "" {
		return errors.New("plant ID and department ID are required")
	}
	if a.Name == "" {
		return errors.New("aspect name is required")
	}
	return nil
}

// EvaluateRiskScore calculates risk value.
func (i *EnvironmentalImpact) EvaluateRiskScore() {
	if i.Severity < 1 {
		i.Severity = 1
	} else if i.Severity > 5 {
		i.Severity = 5
	}
	if i.Probability < 1 {
		i.Probability = 1
	} else if i.Probability > 5 {
		i.Probability = 5
	}
	i.RiskScore = i.Severity * i.Probability
}
