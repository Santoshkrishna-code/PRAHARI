package ghsclassification

// GHS represents hazard classification elements mapped to a chemical.
type GHS struct {
	ID          string `json:"id"`
	ChemicalID  string `json:"chemical_id"`
	SignalWord  string `json:"signal_word"` // DANGER, WARNING, NONE
	HazardClass string `json:"hazard_class"` // FLAMMABLE, TOXIC, CORROSIVE, etc.
	Category    string `json:"category"`     // Category 1, 2, 3
}
