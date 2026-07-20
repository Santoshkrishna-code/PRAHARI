package downtime

import (
	"errors"
	"time"
)

// Downtime tracks production interruption.
type Downtime struct {
	ID           string    `json:"id" db:"id"`
	MaintenanceID string   `json:"maintenance_id" db:"maintenance_id"`
	AssetID      string    `json:"asset_id" db:"asset_id"`
	StartDate    time.Time `json:"start_date" db:"start_date"`
	EndDate      time.Time `json:"end_date" db:"end_date"`
	ReasonCode   string    `json:"reason_code" db:"reason_code"` // Breakdown, Planned, etc
}

// Validate checks domain invariants.
func (d *Downtime) Validate() error {
	if d.MaintenanceID == "" {
		return errors.New("maintenance ID is required")
	}
	if d.AssetID == "" {
		return errors.New("asset ID reference is required")
	}
	if d.EndDate.Before(d.StartDate) {
		return errors.New("end date must exceed downtime start date")
	}
	return nil
}

// DowntimeMinutes calculates total duration.
func (d *Downtime) DowntimeMinutes() float64 {
	return d.EndDate.Sub(d.StartDate).Minutes()
}
