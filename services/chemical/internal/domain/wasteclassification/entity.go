package wasteclassification

// Classification defines hazardous waste categorization criteria (RCRA, EWC, etc.).
type Classification struct {
	ID          string `json:"id"`
	ChemicalID  string `json:"chemical_id"`
	RCRACode    string `json:"rcra_code,omitempty"` // E.g., D001, U002
	EWCCode     string `json:"ewc_code,omitempty"`  // European Waste Catalogue code
	HazardClass string `json:"hazard_class"`        // IGNITABLE, CORROSIVE, REACTIVE, TOXIC
}
