package events

import (
	"time"
)

// ProfileCreated event.
type ProfileCreated struct {
	ID           string    `json:"id"`
	FacilityName string    `json:"facility_name"`
	Timestamp    time.Time `json:"timestamp"`
}

// ReadingRecorded event.
type ReadingRecorded struct {
	ID           string    `json:"id"`
	ReadingValue float64   `json:"reading_value"`
	Timestamp    time.Time `json:"timestamp"`
}

// ForecastGenerated event.
type ForecastGenerated struct {
	ID             string    `json:"id"`
	PredictedKWh   float64   `json:"predicted_kwh"`
	ConfidenceRate float64   `json:"confidence_rate"`
	Timestamp      time.Time `json:"timestamp"`
}

// TargetExceeded event.
type TargetExceeded struct {
	ID             string    `json:"id"`
	PlantID        string    `json:"plant_id"`
	TriggerMessage string    `json:"trigger_message"`
	Timestamp      time.Time `json:"timestamp"`
}

// OptimizationRecommended event.
type OptimizationRecommended struct {
	ID           string    `json:"id"`
	EstSavingUSD float64   `json:"est_saving_usd"`
	Timestamp    time.Time `json:"timestamp"`
}

// RenewableGenerated event.
type RenewableGenerated struct {
	ID           string    `json:"id"`
	KWhGenerated float64   `json:"kwh_generated"`
	Timestamp    time.Time `json:"timestamp"`
}
