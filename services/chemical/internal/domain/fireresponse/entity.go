package fireresponse

// Guidance holds instructions for extinguishing fires involving specific chemicals.
type Guidance struct {
	ID             string `json:"id"`
	ChemicalID     string `json:"chemical_id"`
	MediaSuitable  string `json:"media_suitable"`   // E.g., CO2, Dry Powder
	MediaUnsuitable string `json:"media_unsuitable"` // E.g., Do not use water!
	HazardProducts string `json:"hazard_products"`  // Toxic combustion products
	FireFighterPPE string `json:"fire_fighter_ppe"`
}
