package firstaid

// Guidance holds instructions for first aid treatment in case of chemical exposure.
type Guidance struct {
	ID         string `json:"id"`
	ChemicalID string `json:"chemical_id"`
	Inhalation string `json:"inhalation"`
	Skin       string `json:"skin"`
	Eyes       string `json:"eyes"`
	Ingestion  string `json:"ingestion"`
}
