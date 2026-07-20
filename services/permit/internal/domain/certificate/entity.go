package certificate

import (
	"encoding/json"
	"errors"
	"time"
)

// Type classifies digital certificates issued at key lifecycle stages.
type Type string

const (
	TypePermitIssued    Type = "PERMIT_ISSUED"
	TypePermitReceived  Type = "PERMIT_RECEIVED"
	TypeWorkCompleted   Type = "WORK_COMPLETED"
	TypePermitClosed    Type = "PERMIT_CLOSED"
)

// Certificate represents an immutable record certifying authorization steps.
type Certificate struct {
	ID              string          `json:"id" db:"id"`
	PermitID        string          `json:"permit_id" db:"permit_id"`
	Type            Type            `json:"type" db:"type"`
	IssuedTo        string          `json:"issued_to" db:"issued_to"`
	IssuedBy        string          `json:"issued_by" db:"issued_by"`
	SignatureHash   string          `json:"signature_hash" db:"signature_hash"`
	IssuedAt        time.Time       `json:"issued_at" db:"issued_at"`
	ExpiresAt       *time.Time      `json:"expires_at,omitempty" db:"expires_at"`
	RevokedAt       *time.Time      `json:"revoked_at,omitempty" db:"revoked_at"`
	CertificateData json.RawMessage `json:"certificate_data" db:"certificate_data"` // JSON snapshot of state
}

// Validate checks domain invariants for Certificate.
func (c *Certificate) Validate() error {
	if c.PermitID == "" {
		return errors.New("permit ID is required for certificate")
	}
	if c.SignatureHash == "" {
		return errors.New("digital signature hash is required")
	}
	if c.IssuedTo == "" {
		return errors.New("recipient identifier is required")
	}
	return nil
}

// Revoke registers revocation details on this certificate.
func (c *Certificate) Revoke() {
	now := time.Now()
	c.RevokedAt = &now
}

// IsRevoked checks revocation state.
func (c *Certificate) IsRevoked() bool {
	return c.RevokedAt != nil
}
