package renewablegeneration

import (
	"errors"
	"time"
)

// Generation tracks clean power generated offsite or onsite.
type Generation struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	SourceType     string    `json:"source_type" db:"source_type"` // "SOLAR", "WIND"
	KWhGenerated   float64   `json:"kwh_generated" db:"kwh_generated"`
	Co2OffsetKg    float64   `json:"co2_offset_kg" db:"co2_offset_kg"`
	RecordedAt     time.Time `json:"recorded_at" db:"recorded_at"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Validate checks fields.
func (g *Generation) Validate() error {
	if g.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	if g.KWhGenerated <= 0.0 {
		return errors.New("generation KWh value must be positive")
	}
	return nil
}
