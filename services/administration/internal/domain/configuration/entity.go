package configuration

import "time"

// Param represents a global or tenant specific configuration parameter value.
type Param struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id,omitempty"`
	ConfigKey string    `json:"config_key"`
	Val       string    `json:"val"`
	UpdatedAt time.Time `json:"updated_at"`
}
