package observationtype

import (
	"errors"
)

// ObservationType classifies safe/unsafe behaviors.
type ObservationType struct {
	ID          string `json:"id" db:"id"`
	Code        string `json:"code" db:"code"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (ot *ObservationType) Validate() error {
	if ot.Code == "" {
		return errors.New("type code is required")
	}
	if ot.Name == "" {
		return errors.New("type name is required")
	}
	return nil
}
