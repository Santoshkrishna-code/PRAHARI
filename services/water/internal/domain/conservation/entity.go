package conservation

import "time"

// Program represents a water conservation initiative.
type Program struct {
	ID              string    `json:"id"`
	PlantID         string    `json:"plant_id"`
	ProgramName     string    `json:"program_name"`
	BaselineKLD     float64   `json:"baseline_kld"`
	TargetSavedKLD  float64   `json:"target_saved_kld"`
	ActualSavedKLD  float64   `json:"actual_saved_kld"`
	InvestmentUSD   float64   `json:"investment_usd"`
	ROIPercent      float64   `json:"roi_percent"`
	Status          string    `json:"status"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
