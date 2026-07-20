package isolationplan

import "time"

// Plan represents an approved engineering sequence/steps to isolate hazardous energy.
type Plan struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	EquipmentID string    `json:"equipment_id"` // Target Asset/System being isolated
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ApprovedBy  string    `json:"approved_by,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
