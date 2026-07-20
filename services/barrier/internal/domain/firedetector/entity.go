package firedetector

import "time"

// Detector represents an optical, flame, or thermal fire detection barrier.
type Detector struct {
	ID           string    `json:"id"`
	BarrierID    string    `json:"barrier_id"`
	TagNumber    string    `json:"tag_number"`
	DetectorType string    `json:"detector_type"` // UV/IR, HEAT, SMOKE
	LocationCode string    `json:"location_code"`
	CreatedAt    time.Time `json:"created_at"`
}
