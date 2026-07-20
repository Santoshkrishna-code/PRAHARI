package pperequirement

import "time"

// Requirement defines hazard-based or task-based mandatory PPE requirements (tied to Risk Assessments or permits).
type Requirement struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	HazardType  string    `json:"hazard_type"` // CHEMICAL, ELECTRICAL, RADIOLOGICAL, HEIGHTS
	WorkArea    string    `json:"work_area"`
	PPEIDList   string    `json:"ppe_id_list"` // Comma-separated list of mandatory PPE IDs
	CreatedAt   time.Time `json:"created_at"`
}
