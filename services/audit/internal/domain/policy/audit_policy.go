package policy

import (
	"errors"

	"prahari/services/audit/internal/domain/audit"
)

// CheckEvidenceRequirements checks if audits require statutory files uploads.
func CheckEvidenceRequirements(a *audit.Audit, evidenceCount int) error {
	if a.StatusCode == "REVIEW" && evidenceCount == 0 {
		return errors.New("audit review requests require at least one verified document file evidence attachment")
	}
	return nil
}
