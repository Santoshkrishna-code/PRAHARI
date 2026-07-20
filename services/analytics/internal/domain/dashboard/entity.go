package dashboard

import "time"

// Dashboard represents an analytical panel definition containing collections of metrics and KPIs.
type Dashboard struct {
	ID        string    `json:"id"`
	PlantID   string    `json:"plant_id"`
	Name      string    `json:"name"` // E.g., Executive HSE Dashboard
	Config    string    `json:"config"` // JSON string config layout
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
}
