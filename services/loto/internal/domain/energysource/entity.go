package energysource

// Source represents an identified hazardous energy source (e.g. 480V Electrical, High Pressure Steam, 100 PSI Air).
type Source struct {
	ID          string `json:"id"`
	PlantID     string `json:"plant_id"`
	Name        string `json:"name"`
	EnergyType  string `json:"energy_type"` // ELECTRICAL, MECHANICAL, HYDRAULIC, PNEUMATIC, THERMAL, CHEMICAL
	Magnitude   string `json:"magnitude"`   // E.g., 415V, 10 bar
	Description string `json:"description"`
}
