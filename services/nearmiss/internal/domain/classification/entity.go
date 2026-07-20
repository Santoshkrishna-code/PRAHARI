package classification

import (
	"errors"
)

// Classification classifies Unsafe Acts, Unsafe Conditions, Process Deviations.
type Classification struct {
	ID          string `json:"id" db:"id"`
	Code        string `json:"code" db:"code"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (c *Classification) Validate() error {
	if c.Code == "" {
		return errors.New("classification code is required")
	}
	if c.Name == "" {
		return errors.New("classification name is required")
	}
	return nil
}
