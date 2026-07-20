package maintenancetype

import (
	"errors"
)

// MaintenanceType classifies preventive, predictive, and breakdown strategies.
type MaintenanceType struct {
	ID          string `json:"id" db:"id"`
	Code        string `json:"code" db:"code"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (mt *MaintenanceType) Validate() error {
	if mt.Code == "" {
		return errors.New("type code is required")
	}
	if mt.Name == "" {
		return errors.New("type name is required")
	}
	return nil
}
