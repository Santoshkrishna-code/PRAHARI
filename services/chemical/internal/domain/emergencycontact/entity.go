package emergencycontact

// Contact holds numbers/details of emergency contacts for HAZMAT handling.
type Contact struct {
	ID          string `json:"id"`
	PlantID     string `json:"plant_id"`
	ContactName string `json:"contact_name"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"` // CHEMTREC, Poison Control, Internal Coordinator
}
