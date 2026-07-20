package contamination

import (
	"errors"
	"time"
)

// SoilContamination registers heavy metal or hydrocarbon ground contamination.
type SoilContamination struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	AgentName      string    `json:"agent_name" db:"agent_name"` // "HYDROCARBON", "LEAD", "MERCURY"
	AreaSqMeters   float64   `json:"area_sq_meters" db:"area_sq_meters"`
	DepthCm        float64   `json:"depth_cm" db:"depth_cm"`
	RemediationStatus string `json:"remediation_status" db:"remediation_status"` // "PENDING", "IN_PROGRESS", "CLEANED"
	Notes          string    `json:"notes" db:"notes"`
	DetectedAt     time.Time `json:"detected_at" db:"detected_at"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Validate checks contamination inputs.
func (c *SoilContamination) Validate() error {
	if c.PlantID == "" {
		return errors.New("plant ID is required")
	}
	if c.AgentName == "" {
		return errors.New("contamination agent name is required")
	}
	return nil
}
