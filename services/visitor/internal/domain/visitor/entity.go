package visitor

import "time"

// Visitor represents an individual seeking physical access to a plant.
type Visitor struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Company      string    `json:"company"`
	VisitorType  string    `json:"visitor_type"` // VIP, AUDITOR, CONTRACTOR, GOVT_OFFICIAL, DELIVERY
	IDType       string    `json:"id_type"`       // PASSPORT, DRIVERS_LICENSE, NATIONAL_ID
	IDNumber     string    `json:"id_number"`
	Blacklisted  bool      `json:"blacklisted"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
