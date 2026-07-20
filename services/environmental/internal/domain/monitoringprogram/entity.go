package monitoringprogram

import (
	"errors"
	"time"
)

// MonitoringProgram represents structured compliance monitoring programs.
type MonitoringProgram struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	ProgramType    string    `json:"program_type" db:"program_type"` // e.g. "AIR_EMISSIONS", "WATER_QUALITY"
	Title          string    `json:"title" db:"title"`
	StartDate      time.Time `json:"start_date" db:"start_date"`
	NextSchedule   time.Time `json:"next_schedule" db:"next_schedule"`
	FrequencyDays  int       `json:"frequency_days" db:"frequency_days"`
	Status         string    `json:"status" db:"status"` // "ACTIVE", "PAUSED", "COMPLETED"
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks program values.
func (p *MonitoringProgram) Validate() error {
	if p.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	if p.ProgramType == "" {
		return errors.New("monitoring program type is required")
	}
	return nil
}
