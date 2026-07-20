package optimization

import (
	"errors"
	"time"
)

// Recommendation structures smart energy optimization findings.
type Recommendation struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	AssetID        string    `json:"asset_id" db:"asset_id"`
	Title          string    `json:"title" db:"title"` // e.g. "Shed load during peak hours 2PM-6PM"
	Description    string    `json:"description" db:"description"`
	EstSavingUSD   float64   `json:"est_saving_usd" db:"est_saving_usd"`
	EstSavingKWh   float64   `json:"est_saving_kwh" db:"est_saving_kwh"`
	Priority       string    `json:"priority" db:"priority"` // "HIGH", "MEDIUM", "LOW"
	Status         string    `json:"status" db:"status"`     // "RECOMMENDED", "IMPLEMENTED", "REJECTED"
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Validate checks limits.
func (r *Recommendation) Validate() error {
	if r.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	if r.Title == "" {
		return errors.New("recommendation title is required")
	}
	return nil
}
