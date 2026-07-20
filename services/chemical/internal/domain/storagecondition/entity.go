package storagecondition

// Condition represents safety conditions / requirements for storing specific chemicals.
type Condition struct {
	ID          string  `json:"id"`
	ChemicalID  string  `json:"chemical_id"`
	TempMinC    float64 `json:"temp_min_c"`
	TempMaxC    float64 `json:"temp_max_c"`
	HumidityMax float64 `json:"humidity_max"` // relative percentage
	Lighting    string  `json:"lighting"`     // E.g., DARK, AVOID_DIRECT_SUNLIGHT
	Ventilation string  `json:"ventilation"`  // REQUIRED, HIGH
}
