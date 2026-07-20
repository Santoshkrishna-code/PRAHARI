package status

import (
	"fmt"
)

// Code defines state tags.
type Code string

const (
	CodeDraft      Code = "DRAFT"
	CodePlanned    Code = "PLANNED"
	CodeScheduled  Code = "SCHEDULED"
	CodeInProgress Code = "IN_PROGRESS"
	CodeEvidence   Code = "EVIDENCE_COLLECTION"
	CodeReview     Code = "REVIEW"
	CodeApproved   Code = "APPROVED"
	CodeClosed     Code = "CLOSED"
)

// ValidCodes lists all valid state codes.
var ValidCodes = []Code{
	CodeDraft,
	CodePlanned,
	CodeScheduled,
	CodeInProgress,
	CodeEvidence,
	CodeReview,
	CodeApproved,
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
