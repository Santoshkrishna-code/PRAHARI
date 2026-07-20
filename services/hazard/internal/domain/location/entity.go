package location

import (
	"errors"
)

// Location coordinates hazard mapping references.
type Location struct {
	ID         string `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	FacilityID string `json:"facility_id" db:"facility_id"`
}

// Validate checks domain invariants.
func (l *Location) Validate() error {
	if l.Name == "" {
		return errors.New("location name is required")
	}
	return nil
}
