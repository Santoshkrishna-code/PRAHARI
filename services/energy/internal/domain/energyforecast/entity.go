package energyforecast

import (
	"errors"
	"time"
)

// Forecast maps predicted load curves.
type Forecast struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	ForecastPeriod string    `json:"forecast_period" db:"forecast_period"` // e.g. "AUGUST-2026"
	PredictedKWh   float64   `json:"predicted_kwh" db:"predicted_kwh"`
	ConfidenceRate float64   `json:"confidence_rate" db:"confidence_rate"` // e.g. 94.5
	GeneratedAt    time.Time `json:"generated_at" db:"generated_at"`
}

// Validate checks forecast.
func (f *Forecast) Validate() error {
	if f.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	return nil
}
