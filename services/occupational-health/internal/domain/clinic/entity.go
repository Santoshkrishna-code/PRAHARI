package clinic

import (
	"errors"
	"time"
)

// Clinic registers authorized healthcare centers.
type Clinic struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Address   string    `json:"address" db:"address"`
	ContactNo string    `json:"contact_no" db:"contact_no"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks clinic details.
func (c *Clinic) Validate() error {
	if c.Name == "" {
		return errors.New("clinic name is required")
	}
	return nil
}
