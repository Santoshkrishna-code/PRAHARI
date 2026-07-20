package forecasting

import "time"

// Forecast represents a predicted water demand.
type Forecast struct {
	ID              string    `json:"id"`
	PlantID         string    `json:"plant_id"`
	ForecastPeriod  string    `json:"forecast_period"`
	PredictedKL     float64   `json:"predicted_kl"`
	ConfidenceRate  float64   `json:"confidence_rate"`
	SeasonalFactor  float64   `json:"seasonal_factor"`
	GeneratedAt     time.Time `json:"generated_at"`
}
