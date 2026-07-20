package calibrationplan

import "time"

// Plan represents a defined plan or procedure for calibrating a class of instruments.
type Plan struct {
	ID             string    `json:"id"`
	PlantID        string    `json:"plant_id"`
	ProcedureName  string    `json:"procedure_name"`
	IntervalMonths int       `json:"interval_months"`
	Instructions   string    `json:"instructions"`
	CreatedAt      time.Time `json:"created_at"`
}
