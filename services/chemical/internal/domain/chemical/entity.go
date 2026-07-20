package chemical

import "time"

// Chemical is the core domain model representing a registered chemical master definition.
type Chemical struct {
	ID               string    `json:"id"`
	PlantID          string    `json:"plant_id"`
	Name             string    `json:"name"`
	CASNumber        string    `json:"cas_number"`
	IUPACName        string    `json:"iupac_name"`
	Formula          string    `json:"formula"`
	MolecularWeight  float64   `json:"molecular_weight"`
	PhysicalState    string    `json:"physical_state"` // SOLID, LIQUID, GAS
	IsRestricted     bool      `json:"is_restricted"`
	MaxAllowableQty  float64   `json:"max_allowable_qty"` // MAQ in kg
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
