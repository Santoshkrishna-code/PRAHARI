package meterreading

import "time"

// Reading represents a single flow meter measurement.
type Reading struct {
	ID            string    `json:"id"`
	MeterID       string    `json:"meter_id"`
	ReadingValueKL float64  `json:"reading_value_kl"`
	FlowRateKLH   float64   `json:"flow_rate_klh"`
	PressureBar   float64   `json:"pressure_bar"`
	TemperatureC  float64   `json:"temperature_c"`
	ReadingTime   time.Time `json:"reading_time"`
	CreatedAt     time.Time `json:"created_at"`
}
