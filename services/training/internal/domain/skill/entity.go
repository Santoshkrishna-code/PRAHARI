package skill

import (
	"errors"
)

// Skill details safety task attributes.
type Skill struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (s *Skill) Validate() error {
	if s.Name == "" {
		return errors.New("skill name is required")
	}
	return nil
}
