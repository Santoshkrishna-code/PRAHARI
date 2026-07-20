package organization

import (
	"errors"
)

// Organization represents tenant profile aggregates.
type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Plan string `json:"plan"`
}

// Validate checks model parameters.
func (o *Organization) Validate() error {
	if o.ID == "" {
		return errors.New("organization ID is required")
	}
	if o.Name == "" {
		return errors.New("organization name is required")
	}
	return nil
}
