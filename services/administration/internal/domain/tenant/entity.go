package tenant

import "time"

// Tenant represents a multi-tenant customer account in the SaaS platform.
type Tenant struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Domain    string    `json:"domain"`
	Status    string    `json:"status"` // ACTIVE, SUSPENDED, TRIAL
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
