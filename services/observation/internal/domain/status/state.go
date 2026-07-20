package status

import (
	"fmt"
)

// Code defines state tags.
type Code string

const (
	CodeObserved Code = "OBSERVED"
	CodeRecorded Code = "RECORDED"
	CodeReviewed Code = "REVIEWED"
	CodeCoaching Code = "COACHING"
	CodeFollowUp Code = "FOLLOWUP"
	CodeVerified Code = "VERIFIED"
	CodeClosed   Code = "CLOSED"
)

// ValidCodes lists all valid state codes.
var ValidCodes = []Code{
	CodeObserved,
	CodeRecorded,
	CodeReviewed,
	CodeCoaching,
	CodeFollowUp,
	CodeVerified,
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
