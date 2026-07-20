package certification

import (
	"errors"
	"time"
)

// Certification maps professional signoffs validation.
type Certification struct {
	ID             string    `json:"id" db:"id"`
	WorkerID       string    `json:"worker_id" db:"worker_id"`
	CertNumber     string    `json:"cert_number" db:"cert_number"`
	Title          string    `json:"title" db:"title"`
	ExpiryDate     time.Time `json:"expiry_date" db:"expiry_date"`
	Issuer         string    `json:"issuer" db:"issuer"`
}

// Validate checks domain invariants.
func (c *Certification) Validate() error {
	if c.WorkerID == "" {
		return errors.New("worker ID reference is required")
	}
	if c.CertNumber == "" {
		return errors.New("certification credentials code is required")
	}
	return nil
}

// IsExpired checks validity periods limits.
func (c *Certification) IsExpired() bool {
	return time.Now().After(c.ExpiryDate)
}
