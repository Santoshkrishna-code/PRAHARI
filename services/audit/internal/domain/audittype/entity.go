package audittype

import (
	"errors"
)

// AuditType classifies internal/external audits.
type AuditType struct {
	ID          string `json:"id" db:"id"`
	Code        string `json:"code" db:"code"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (at *AuditType) Validate() error {
	if at.Code == "" {
		return errors.New("audit type code is required")
	}
	if at.Name == "" {
		return errors.New("audit type name is required")
	}
	return nil
}
