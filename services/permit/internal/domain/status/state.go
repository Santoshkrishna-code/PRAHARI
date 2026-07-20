package status

import (
	"fmt"
)

// Code defines state tags for a permit.
type Code string

const (
	CodeDraft          Code = "DRAFT"
	CodeSubmitted      Code = "SUBMITTED"
	CodeRiskAssessment Code = "RISK_ASSESSMENT"
	CodeApproval       Code = "APPROVAL"
	CodeIssued         Code = "ISSUED"
	CodeActive         Code = "ACTIVE"
	CodeSuspended      Code = "SUSPENDED"
	CodeCompleted      Code = "COMPLETED"
	CodeClosed         Code = "CLOSED"
	CodeArchived       Code = "ARCHIVED"
	CodeRejected       Code = "REJECTED"
	CodeCancelled      Code = "CANCELLED"
)

// ValidCodes lists all valid state codes.
var ValidCodes = []Code{
	CodeDraft,
	CodeSubmitted,
	CodeRiskAssessment,
	CodeApproval,
	CodeIssued,
	CodeActive,
	CodeSuspended,
	CodeCompleted,
	CodeClosed,
	CodeArchived,
	CodeRejected,
	CodeCancelled,
}

// Validate checks if status exists.
func (c Code) Validate() error {
	for _, valid := range ValidCodes {
		if c == valid {
			return nil
		}
	}
	return fmt.Errorf("invalid status code: %s", c)
}

// IsTerminal checks if the status is a final state.
func (c Code) IsTerminal() bool {
	return c == CodeClosed || c == CodeArchived || c == CodeRejected || c == CodeCancelled
}
