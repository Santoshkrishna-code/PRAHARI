package status

import (
	"fmt"
)

// Code defines state tags.
type Code string

const (
	CodeDraft      Code = "DRAFT"
	CodeAssessment Code = "ASSESSMENT"
	CodeReview     Code = "REVIEW"
	CodeApproval   Code = "APPROVAL"
	CodeActive     Code = "ACTIVE"
	CodeReassess   Code = "REASSESSMENT"
	CodeClosed     Code = "CLOSED"
)

// ValidCodes lists all valid state codes.
var ValidCodes = []Code{
	CodeDraft,
	CodeAssessment,
	CodeReview,
	CodeApproval,
	CodeActive,
	CodeReassess,
	CodeClosed,
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
