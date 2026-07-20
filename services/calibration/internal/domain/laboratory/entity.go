package laboratory

import "time"

// Laboratory represents an internal or external metrology lab coordinates and ISO/IEC 17025 accreditations.
type Laboratory struct {
	ID             string    `json:"id"`
	LabName        string    `json:"lab_name"`
	Accreditation  string    `json:"accreditation"` // E.g. ISO/IEC 17025, A2LA
	ContactPerson  string    `json:"contact_person"`
	Email          string    `json:"email"`
	Active         bool      `json:"active"`
	CreatedAt      time.Time `json:"created_at"`
}
