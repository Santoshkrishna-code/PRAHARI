package precautionstatement

// Statement represents a GHS precautionary statement (e.g. P102, P210).
type Statement struct {
	ID         string `json:"id"`
	ChemicalID string `json:"chemical_id"`
	PCode      string `json:"p_code"`      // E.g., P210
	Statement  string `json:"statement"`   // E.g., Keep away from heat...
}
