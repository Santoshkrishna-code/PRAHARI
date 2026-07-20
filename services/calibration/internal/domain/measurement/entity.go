package measurement

import "time"

// Result records measurement test points (nominal value, standard value, as-found value, as-left value).
type Result struct {
	ID            string    `json:"id"`
	CalibrationID string    `json:"calibration_id"`
	TestPoint     float64   `json:"test_point"`
	NominalValue  float64   `json:"nominal_value"`
	StandardValue float64   `json:"standard_value"`
	AsFoundValue  float64   `json:"as_found_value"`
	AsLeftValue   float64   `json:"as_left_value"`
	Timestamp     time.Time `json:"timestamp"`
}
