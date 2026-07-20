package source

import (
	"errors"
)

// Source classifies where a hazard report originated (Inspection, Incident, Observation).
type Source struct {
	ID          string `json:"id" db:"id"`
	Code        string `json:"code" db:"code"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (s *Source) Validate() error {
	if s.Code == "" {
		return errors.New("source code is required")
	}
	return nil
}
