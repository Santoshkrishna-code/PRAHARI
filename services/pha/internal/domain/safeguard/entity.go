package safeguard

import "time"

// Safeguard represents existing protective controls or Independent Protection Layers (IPL).
type Safeguard struct {
	ID           string    `json:"id"`
	ScenarioID   string    `json:"scenario_id"`
	Title        string    `json:"title"`
	SafeguardType string   `json:"safeguard_type"` // ENG_CONTROL, ADMIN_CONTROL, SIS, ALARM, RELIEF_VALVE
	IsIPL        bool      `json:"is_ipl"`         // Independent Protection Layer
	PFD          float64   `json:"pfd"`            // Probability of Failure on Demand for LOPA
	AssetID      string    `json:"asset_id,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}
