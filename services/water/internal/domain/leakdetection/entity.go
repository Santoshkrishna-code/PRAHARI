package leakdetection

import "time"

// Leak represents a detected water leakage event.
type Leak struct {
	ID              string    `json:"id"`
	PlantID         string    `json:"plant_id"`
	PipelineID      string    `json:"pipeline_id,omitempty"`
	ZoneCode        string    `json:"zone_code"`
	DetectionMethod string    `json:"detection_method"`
	EstimatedLossKLD float64  `json:"estimated_loss_kld"`
	Severity        string    `json:"severity"`
	LocationDesc    string    `json:"location_desc"`
	WorkOrderID     string    `json:"work_order_id,omitempty"`
	IsResolved      bool      `json:"is_resolved"`
	DetectedAt      time.Time `json:"detected_at"`
	ResolvedAt      time.Time `json:"resolved_at,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
}
