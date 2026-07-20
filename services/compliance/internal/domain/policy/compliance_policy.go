package policy

import (
	"errors"

	"prahari/services/compliance/internal/domain/compliance"
)

// CheckEvidenceRequirements checks if obligations require statutory files uploads.
func CheckEvidenceRequirements(c *compliance.Compliance, evidenceCount int) error {
	if c.StatusCode == "REVIEW" && evidenceCount == 0 {
		return errors.New("compliance review requests require at least one verified document file evidence attachment")
	}
	return nil
}
