package recycling

import "time"

// Program represents a water recycling initiative.
type Program struct {
	ID              string    `json:"id"`
	PlantID         string    `json:"plant_id"`
	ProgramName     string    `json:"program_name"`
	SourceType      string    `json:"source_type"`
	TreatmentMethod string    `json:"treatment_method"`
	InputKLD        float64   `json:"input_kld"`
	OutputKLD       float64   `json:"output_kld"`
	RecycleRatePct  float64   `json:"recycle_rate_pct"`
	Status          string    `json:"status"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
