package status

import (
	"fmt"
)

// Code defines state tags.
type Code string

const (
	CodeIdentified  Code = "IDENTIFIED"
	CodeReported    Code = "REPORTED"
	CodeAssessed    Code = "RISK_ASSESSMENT"
	CodeMitigating  Code = "MITIGATION_PLANNING"
	CodeApproved    Code = "APPROVAL"
	CodeImplemented Code = "IMPLEMENTATION"
	CodeVerified    Code = "VERIFICATION"
	CodeClosed      Code = "CLOSED"
	CodeRejected    Code = "REJECTED"
)

// ValidCodes lists all valid state codes.
var ValidCodes = []Code{
	CodeIdentified,
	CodeReported,
	CodeAssessed,
	CodeMitigating,
	CodeApproved,
	CodeImplemented,
	CodeVerified,
	CodeClosed,
	CodeRejected,
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
