package status

import (
	"fmt"
)

// Code defines state tags in inspections.
type Code string

const (
	CodeDraft       Code = "DRAFT"
	CodeScheduled   Code = "SCHEDULED"
	CodeAssigned    Code = "ASSIGNED"
	CodeInProgress  Code = "IN_PROGRESS"
	CodeCompleted   Code = "COMPLETED"
	CodeUnderReview Code = "UNDER_REVIEW"
	CodeApproved    Code = "APPROVED"
	CodeClosed      Code = "CLOSED"
	CodeArchived    Code = "ARCHIVED"
	CodeRejected    Code = "REJECTED"
	CodeCancelled   Code = "CANCELLED"
)

// ValidCodes lists all valid state codes.
var ValidCodes = []Code{
	CodeDraft,
	CodeScheduled,
	CodeAssigned,
	CodeInProgress,
	CodeCompleted,
	CodeUnderReview,
	CodeApproved,
	CodeClosed,
	CodeArchived,
	CodeRejected,
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
