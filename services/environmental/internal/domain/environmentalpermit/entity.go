package environmentalpermit

import (
	"errors"
	"time"
)

// EnvironmentalPermit tracks regulatory operating licenses.
type EnvironmentalPermit struct {
	ID             string    `json:"id" db:"id"`
	PermitNumber   string    `json:"permit_number" db:"permit_number"`
	Title          string    `json:"title" db:"title"`
	Agency         string    `json:"agency" db:"agency"` // e.g. "EPA", "PCB"
	IssueDate      time.Time `json:"issue_date" db:"issue_date"`
	ExpiryDate     time.Time `json:"expiry_date" db:"expiry_date"`
	Status         string    `json:"status" db:"status"` // "ACTIVE", "EXPIRED", "SUSPENDED"
	ConditionsText string    `json:"conditions_text" db:"conditions_text"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks permit constraints.
func (p *EnvironmentalPermit) Validate() error {
	if p.PermitNumber == "" {
		return errors.New("permit number is required")
	}
	if p.Title == "" {
		return errors.New("permit title is required")
	}
	if p.ExpiryDate.Before(p.IssueDate) {
		return errors.New("expiry date must be after issue date")
	}
	return nil
}
