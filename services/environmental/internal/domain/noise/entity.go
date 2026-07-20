package noise

import (
	"errors"
	"time"
)

// NoiseMonitoring records sound levels at perimeter locations.
type NoiseMonitoring struct {
	ID             string    `json:"id" db:"id"`
	LocationID     string    `json:"location_id" db:"location_id"`
	DecibelsDbA    float64   `json:"decibels_dba" db:"decibels_dba"`
	DurationMins   int       `json:"duration_mins" db:"duration_mins"`
	LimitThreshold float64   `json:"limit_threshold" db:"limit_threshold"`
	IsOverLimit     bool      `json:"is_over_limit" db:"is_over_limit"`
	RecordedAt     time.Time `json:"recorded_at" db:"recorded_at"`
}

// Validate checks noise records.
func (n *NoiseMonitoring) Validate() error {
	if n.LocationID == "" {
		return errors.New("monitoring location ID is required")
	}
	return nil
}
