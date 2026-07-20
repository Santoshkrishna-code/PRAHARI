package waterprofile

import "time"

// Profile represents a facility-level water management profile.
type Profile struct {
	ID               string    `json:"id"`
	PlantID          string    `json:"plant_id"`
	DepartmentID     string    `json:"department_id,omitempty"`
	FacilityName     string    `json:"facility_name"`
	WaterBasinRegion string    `json:"water_basin_region"`
	AnnualBudgetKL   float64   `json:"annual_budget_kl"`
	TargetRecyclePct float64   `json:"target_recycle_pct"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
