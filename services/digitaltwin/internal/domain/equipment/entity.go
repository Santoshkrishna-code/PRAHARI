package equipment

// Asset represents an equipment node in the digital twin graph.
type Asset struct {
	ID         string `json:"id"`
	TwinID     string `json:"twin_id"`
	ExternalID string `json:"external_id"` // Reference to Asset service ID
	Tag        string `json:"tag"` // E.g., P-101A
	AssetType  string `json:"asset_type"` // PUMP, VALVE, VESSEL, MOTOR, SENSOR
}
