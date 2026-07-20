package status

import (
	"fmt"
)

// Code defines state tags.
type Code string

const (
	CodeReported    Code = "REPORTED"
	CodeClassified  Code = "CLASSIFIED"
	CodeInvestigate Code = "INVESTIGATION"
	CodeRootCause   Code = "ROOT_CAUSE_ANALYSIS"
	CodeCorrective  Code = "CORRECTIVE_ACTIONS"
	CodeVerified    Code = "VERIFICATION"
	CodeClosed      Code = "CLOSED"
	CodeEscalated   Code = "ESCALATED"
)

// ValidCodes lists all valid state codes.
var ValidCodes = []Code{
	CodeReported,
	CodeClassified,
	CodeInvestigate,
	CodeRootCause,
	CodeCorrective,
	CodeVerified,
	CodeClosed,
	CodeEscalated,
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
