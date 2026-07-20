package regulation

import (
	"errors"
)

// Regulation maps legal/statutory items.
type Regulation struct {
	ID          string `json:"id" db:"id"`
	Code        string `json:"code" db:"code"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (r *Regulation) Validate() error {
	if r.Code == "" {
		return errors.New("regulation code is required")
	}
	if r.Name == "" {
		return errors.New("regulation name is required")
	}
	return nil
}
