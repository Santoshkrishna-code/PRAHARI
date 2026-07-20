package airquality

import (
	"errors"
	"time"
)

// AirQuality records perimeter ambient air quality parameters.
type AirQuality struct {
	ID             string    `json:"id" db:"id"`
	StationID      string    `json:"station_id" db:"station_id"`
	AQI            int       `json:"aqi" db:"aqi"`
	PM10           float64   `json:"pm10" db:"pm10"`
	PM25           float64   `json:"pm25" db:"pm25"`
	NO2            float64   `json:"no2" db:"no2"`
	SO2            float64   `json:"so2" db:"so2"`
	O3             float64   `json:"o3" db:"o3"`
	RecordedAt     time.Time `json:"recorded_at" db:"recorded_at"`
	LimitExceeded  bool      `json:"limit_exceeded" db:"limit_exceeded"`
}

// Validate checks quality details.
func (a *AirQuality) Validate() error {
	if a.StationID == "" {
		return errors.New("monitoring station ID is required")
	}
	return nil
}
