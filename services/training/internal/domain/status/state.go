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
	CodeEnrollment Code = "ENROLLMENT"
	CodeInProgress Code = "IN_PROGRESS"
	CodeAssessment Code = "ASSESSMENT"
	CodeCertified  Code = "CERTIFIED"
	CodeActive     Code = "ACTIVE"
	CodeRenewal    Code = "RENEWAL"
)

// ValidCodes lists all valid state codes.
var ValidCodes = []Code{
	CodeDraft,
	CodePlanned,
	CodeScheduled,
	CodeEnrollment,
	CodeInProgress,
	CodeAssessment,
	CodeCertified,
	CodeActive,
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
