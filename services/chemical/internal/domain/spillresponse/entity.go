package spillresponse

// Procedure holds actions, equipment, and steps required to contain a chemical spill safely.
type Procedure struct {
	ID          string `json:"id"`
	ChemicalID  string `json:"chemical_id"`
	SpillKitReq string `json:"spill_kit_req"` // E.g., Acid Spill Kit, Neutraliser
	Containment string `json:"containment"`   // Containment procedure text
	Absorbent   string `json:"absorbent"`     // Absorbent material guidelines
	PPERequired string `json:"ppe_required"`
}
