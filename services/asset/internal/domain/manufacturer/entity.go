package manufacturer

import (
	"errors"
)

// Manufacturer tracks contact details.
type Manufacturer struct {
	ID             string `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	ContactEmail   string `json:"contact_email" db:"contact_email"`
	ContactPhone   string `json:"contact_phone" db:"contact_phone"`
	SupportWebsite string `json:"support_website" db:"support_website"`
}

// Validate checks domain invariants.
func (m *Manufacturer) Validate() error {
	if m.Name == "" {
		return errors.New("manufacturer name is required")
	}
	return nil
}
