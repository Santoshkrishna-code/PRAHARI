package emission

import (
	"errors"
	"time"
)

// Emission represents continuous or point-source atmospheric release metrics.
type Emission struct {
	ID             string    `json:"id" db:"id"`
	SourceID       string    `json:"source_id" db:"source_id"` // Asset ID references e.g. stacks, exhaust
	GasType        string    `json:"gas_type" db:"gas_type"`   // "CO2", "SOx", "NOx", "CO", "PM2.5"
	ReleaseRate    float64   `json:"release_rate" db:"release_rate"`
	UnitOfMeasure  string    `json:"unit_of_measure" db:"unit_of_measure"` // "KG_HR", "PPM", "MG_M3"
	LimitThreshold float64   `json:"limit_threshold" db:"limit_threshold"`
	IsExceeded     bool      `json:"is_exceeded" db:"is_exceeded"`
	ReadingTime    time.Time `json:"reading_time" db:"reading_time"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Validate checks emission inputs.
func (e *Emission) Validate() error {
	if e.SourceID == "" {
		return errors.New("emission source ID is required")
	}
	if e.GasType == "" {
		return errors.New("gas type classification is required")
	}
	return nil
}
