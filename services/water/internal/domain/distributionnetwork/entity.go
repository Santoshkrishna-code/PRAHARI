package distributionnetwork

import "time"

// Network represents a water distribution zone within a facility.
type Network struct {
	ID            string    `json:"id"`
	PlantID       string    `json:"plant_id"`
	ZoneName      string    `json:"zone_name"`
	ZoneCode      string    `json:"zone_code"`
	SupplySourceID string   `json:"supply_source_id"`
	DesignFlowKLD float64   `json:"design_flow_kld"`
	ActualFlowKLD float64   `json:"actual_flow_kld"`
	LossPercent   float64   `json:"loss_percent"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
