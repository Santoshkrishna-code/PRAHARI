package auditprogram

import (
	"errors"
)

// AuditProgram defines annual scheduling tracks.
type AuditProgram struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (ap *AuditProgram) Validate() error {
	if ap.Name == "" {
		return errors.New("program name is required")
	}
	return nil
}
