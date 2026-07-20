package competency

import (
	"errors"
)

// Competency defines safety capability metrics parameters.
type Competency struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (c *Competency) Validate() error {
	if c.Name == "" {
		return errors.New("competency name is required")
	}
	return nil
}
