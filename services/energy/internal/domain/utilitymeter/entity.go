package utilitymeter

import (
	"errors"
	"time"
)

// Meter monitors smart utility connections.
type Meter struct {
	ID             string    `json:"id" db:"id"`
	MeterNumber    string    `json:"meter_number" db:"meter_number"`
	SourceID       string    `json:"source_id" db:"source_id"` // References energysource.ID
	AssetID        string    `json:"asset_id" db:"asset_id"`   // References asset service
	LocationCode   string    `json:"location_code" db:"location_code"`
	Status         string    `json:"status" db:"status"` // "METER_REGISTERED", "DATA_COLLECTION", "ARCHIVED"
	LastCalibrated time.Time `json:"last_calibrated" db:"last_calibrated"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks meter values.
func (m *Meter) Validate() error {
	if m.MeterNumber == "" {
		return errors.New("meter serial number code is required")
	}
	if m.SourceID == "" {
		return errors.New("parent energy source reference ID is required")
	}
	return nil
}
