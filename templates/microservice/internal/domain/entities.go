package domain

import (
	"errors"
	"time"
)

// Incident represents the core DDD aggregate entity of the service.
type Incident struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// Validate checks entity integrity rules.
func (i *Incident) Validate() error {
	if i.ID == "" {
		return errors.New("incident ID cannot be empty")
	}
	if i.Title == "" {
		return errors.New("incident title cannot be empty")
	}
	return nil
}
