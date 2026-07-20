package license

import "time"

// License represents a tenant licensing model quota and expiration definition.
type License struct {
	ID         string    `json:"id"`
	TenantID   string    `json:"tenant_id"`
	Tier       string    `json:"tier"` // ENTERPRISE, PROFESSIONAL, BASIC
	MaxPlants  int       `json:"max_plants"`
	MaxUsers   int       `json:"max_users"`
	ExpiresAt  time.Time `json:"expires_at"`
	LicensedAt time.Time `json:"licensed_at"`
}
