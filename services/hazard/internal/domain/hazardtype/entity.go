package hazardtype

import (
	"errors"
)

// HazardType classifies physical, mechanical, and ergonomic safety bounds.
type HazardType struct {
	ID          string `json:"id" db:"id"`
	Code        string `json:"code" db:"code"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (ht *HazardType) Validate() error {
	if ht.Code == "" {
		return errors.New("type code is required")
	}
	if ht.Name == "" {
		return errors.New("type name is required")
	}
	return nil
}
