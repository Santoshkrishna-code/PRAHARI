package status

import (
	"fmt"
)

// Code defines state tags.
type Code string

const (
	CodeDraft      Code = "DRAFT"
	CodeDefined    Code = "DEFINED"
	CodeAssigned   Code = "ASSIGNED"
	CodeEvidence   Code = "EVIDENCE_COLLECTION"
	CodeReview     Code = "REVIEW"
	CodeCompliant  Code = "COMPLIANT"
	CodeMonitoring Code = "MONITORING"
	CodeRenewal    Code = "RENEWAL"
)

// ValidCodes lists all valid state codes.
var ValidCodes = []Code{
	CodeDraft,
	CodeDefined,
	CodeAssigned,
	CodeEvidence,
	CodeReview,
	CodeCompliant,
	CodeMonitoring,
	CodeRenewal,
}

// Validate checks if code exists.
func (c Code) Validate() error {
	for _, valid := range ValidCodes {
		if c == valid {
			return nil
		}
	}
	return fmt.Errorf("invalid status code: %s", c)
}
