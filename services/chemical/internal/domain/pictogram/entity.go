package pictogram

// Pictogram represents a GHS hazard pictogram.
type Pictogram struct {
	ID          string `json:"id"`
	ChemicalID  string `json:"chemical_id"`
	Code        string `json:"code"`      // GHS01, GHS02, etc.
	Name        string `json:"name"`      // Flame, Corrosion, Exclamation Mark, etc.
	ImageURL    string `json:"image_url"`
}
