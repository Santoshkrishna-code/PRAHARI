package flowmeter

import "time"

// FlowMeter represents a water flow measurement device.
type FlowMeter struct {
	ID              string    `json:"id"`
	PlantID         string    `json:"plant_id"`
	MeterNumber     string    `json:"meter_number"`
	MeterType       string    `json:"meter_type"`
	SourceID        string    `json:"source_id,omitempty"`
	PipelineID      string    `json:"pipeline_id,omitempty"`
	AssetID         string    `json:"asset_id,omitempty"`
	LocationCode    string    `json:"location_code"`
	UnitOfMeasure   string    `json:"unit_of_measure"`
	Status          string    `json:"status"`
	LastCalibratedAt time.Time `json:"last_calibrated_at"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
