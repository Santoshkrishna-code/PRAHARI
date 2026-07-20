package organization

import "time"

// Organization represents a corporate entity within a tenant.
type Organization struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	Name      string    `json:"name"`
	LegalName string    `json:"legal_name"`
	TaxID     string    `json:"tax_id"`
	Status    string    `json:"status"` // DRAFT, CONFIGURED, VALIDATED, ACTIVATED, OPERATIONAL
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
