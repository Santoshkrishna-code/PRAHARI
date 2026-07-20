package contractor

import (
	"errors"
	"time"
)

// Contractor is the central aggregate root of the Contractor Management domain.
type Contractor struct {
	ID               string    `json:"id" db:"id"`
	ContractorNumber string    `json:"contractor_number" db:"contractor_number"`
	CompanyName      string    `json:"company_name" db:"company_name"`
	TaxID            string    `json:"tax_id" db:"tax_id"`
	StatusCode       string    `json:"status_code" db:"status_code"`
	DepartmentID     string    `json:"department_id" db:"department_id"`
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
	InsuranceExpiry  time.Time `json:"insurance_expiry" db:"insurance_expiry"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	IsDeleted        bool      `json:"is_deleted" db:"is_deleted"`
}

// Validate checks domain invariants for Contractor.
func (c *Contractor) Validate() error {
	if c.CompanyName == "" {
		return errors.New("contractor company name is required")
	}
	if len(c.CompanyName) > 200 {
		return errors.New("company name must not exceed 200 characters")
	}
	if c.TaxID == "" {
		return errors.New("tax ID is required")
	}
	if c.DepartmentID == "" {
		return errors.New("department ID is required")
	}
	return nil
}
