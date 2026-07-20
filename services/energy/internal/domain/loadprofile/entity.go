package loadprofile

import (
	"errors"
	"time"
)

// Profile records power demand intervals.
type Profile struct {
	ID             string    `json:"id" db:"id"`
	MeterID        string    `json:"meter_id" db:"meter_id"`
	DemandKW       float64   `json:"demand_kw" db:"demand_kw"`
	TimeInterval   time.Time `json:"time_interval" db:"time_interval"`
	IsPeakPeriod   bool      `json:"is_peak_period" db:"is_peak_period"`
}

// Validate checks profile.
func (p *Profile) Validate() error {
	if p.MeterID == "" {
		return errors.New("meter reference ID is required")
	}
	return nil
}
