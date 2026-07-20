package changetype

import "time"

// TypeDef represents a pre-configured change type definition.
type TypeDef struct {
	ID          string    `json:"id"`
	TypeCode    string    `json:"type_code"` // PROCESS, MECHANICAL, ELECTRICAL, INSTRUMENTATION, SOFTWARE, AUTOMATION, CHEMICAL, ORGANIZATIONAL, DOCUMENTATION, TEMPORARY, EMERGENCY
	TypeName    string    `json:"type_name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}
