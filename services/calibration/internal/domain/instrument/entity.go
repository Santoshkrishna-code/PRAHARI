package instrument

import "time"

// Instrument represents a measuring device or sensor registered in the plant (e.g. pressure transmitter, gas detector).
type Instrument struct {
	ID             string    `json:"id"`
	AssetID        string    `json:"asset_id"` // Link to Asset Management
	PlantID        string    `json:"plant_id"`
	TagNumber      string    `json:"tag_number"` // Tag in P&ID / plant layout
	ModelNumber    string    `json:"model_number"`
	Manufacturer   string    `json:"manufacturer"`
	InstrumentType string    `json:"instrument_type"` // PRESSURE, TEMPERATURE, FLOW, LEVEL, GAS_DETECTOR
	Status         string    `json:"status"`          // ACTIVE, CALIBRATION_DUE, OUT_OF_SERVICE, RETIRED
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
