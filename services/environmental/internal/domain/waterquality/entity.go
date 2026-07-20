package waterquality

import (
	"errors"
	"time"
)

// WaterQuality records groundwater or discharge values.
type WaterQuality struct {
	ID             string    `json:"id" db:"id"`
	LocationID     string    `json:"location_id" db:"location_id"`
	PH             float64   `json:"ph" db:"ph"`
	TurbidityNTU   float64   `json:"turbidity_ntu" db:"turbidity_ntu"`
	DissolvedOxygen float64  `json:"dissolved_oxygen" db:"dissolved_oxygen"` // mg/L
	TDS            float64   `json:"tds" db:"tds"`                           // Total Dissolved Solids mg/L
	Conductivity   float64   `json:"conductivity" db:"conductivity"`         // uS/cm
	TemperatureC   float64   `json:"temperature_c" db:"temperature_c"`
	IsCompliant    bool      `json:"is_compliant" db:"is_compliant"`
	RecordedAt     time.Time `json:"recorded_at" db:"recorded_at"`
}

// Validate checks quality boundaries.
func (w *WaterQuality) Validate() error {
	if w.LocationID == "" {
		return errors.New("monitoring location ID is required")
	}
	if w.PH < 0.0 || w.PH > 14.0 {
		return errors.New("pH level must fall within 0.0 and 14.0 range")
	}
	return nil
}
