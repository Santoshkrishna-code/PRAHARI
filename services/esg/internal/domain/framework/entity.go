package framework

import (
	"errors"
	"time"
)

// Framework registers corporate ESG disclosure guidelines.
type Framework struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"` // "GRI", "SASB", "TCFD", "CDP", "CSRD"
	Description string    `json:"description" db:"description"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Validate checks standard values.
func (f *Framework) Validate() error {
	if f.Name == "" {
		return errors.New("framework compliance name is required")
	}
	return nil
}
