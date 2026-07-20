package hazardstatement

// Statement represents a GHS hazard statement (e.g. H224, H301).
type Statement struct {
	ID         string `json:"id"`
	ChemicalID string `json:"chemical_id"`
	HCode      string `json:"h_code"`      // E.g., H302
	Statement  string `json:"statement"`   // E.g., Harmful if swallowed
}
