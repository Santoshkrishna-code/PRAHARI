package monitoring

import "time"

// Record represents chemical exposure monitoring log.
type Record struct {
	ID          string    `json:"id"`
	ChemicalID  string    `json:"chemical_id"`
	EmployeeID  string    `json:"employee_id"`
	ExposurePPM float64   `json:"exposure_ppm"`
	MeasuredAt  time.Time `json:"measured_at"`
	MeasuredBy  string    `json:"measured_by"`
	Comments    string    `json:"comments,omitempty"`
}
