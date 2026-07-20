package status

import (
	"fmt"
)

// Code defines state tags in walkthrough audits.
type Code string

const (
	CodeDraft       Code = "DRAFT"
	CodePlanned     Code = "PLANNED"
	CodeApproved    Code = "APPROVED"
	CodeScheduled   Code = "SCHEDULED"
	CodeAssigned    Code = "ASSIGNED"
	CodeInProgress  Code = "IN_PROGRESS"
	CodeCompleted   Code = "COMPLETED"
	CodeVerified    Code = "VERIFIED"
	CodeClosed      Code = "CLOSED"
	CodeCancelled   Code = "CANCELLED"
)

// ValidCodes lists all valid state codes.
var ValidCodes = []Code{
	CodeDraft,
	CodePlanned,
	CodeApproved,
	CodeScheduled,
	CodeAssigned,
	CodeInProgress,
	CodeCompleted,
	CodeVerified,
	CodeClosed,
	CodeCancelled,
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
