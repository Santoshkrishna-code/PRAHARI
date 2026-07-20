package mutualaid

import "time"

// Agreement represents a mutual aid covenant with neighboring industrial facilities or municipal fire services.
type Agreement struct {
	ID           string    `json:"id"`
	PlantID      string    `json:"plant_id"`
	PartnerName  string    `json:"partner_name"`
	ContactPhone string    `json:"contact_phone"`
	AidProvided  string    `json:"aid_provided"` // Fire engines, hazmat response, foam supply
	Active       bool      `json:"active"`
	CreatedAt    time.Time `json:"created_at"`
}
