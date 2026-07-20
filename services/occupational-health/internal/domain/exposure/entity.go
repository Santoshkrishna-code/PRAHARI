package exposure

import (
	"errors"
	"time"
)

// ExposureRecord represents monitoring of physical/chemical agent hazards.
type ExposureRecord struct {
	ID              string    `json:"id" db:"id"`
	HealthProfileID string    `json:"health_profile_id" db:"health_profile_id"`
	AgentName       string    `json:"agent_name" db:"agent_name"` // "CHEMICAL", "NOISE", "RADIATION", "ASBESTOS"
	ExposureLevel   float64   `json:"exposure_level" db:"exposure_level"`
	UnitOfMeasure   string    `json:"unit_of_measure" db:"unit_of_measure"` // "PPM", "DBA", "MSV"
	LimitThreshold  float64   `json:"limit_threshold" db:"limit_threshold"`
	MonitoringDate  time.Time `json:"monitoring_date" db:"monitoring_date"`
	IsOverLimit     bool      `json:"is_over_limit" db:"is_over_limit"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks limits and metadata.
func (e *ExposureRecord) Validate() error {
	if e.HealthProfileID == "" {
		return errors.New("health profile reference is required")
	}
	if e.AgentName == "" {
		return errors.New("hazard agent name is required")
	}
	return nil
}
