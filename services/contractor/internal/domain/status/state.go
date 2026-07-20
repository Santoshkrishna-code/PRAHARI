package status

import (
	"fmt"
)

// Code defines state tags.
type Code string

const (
	CodeRegistered    Code = "REGISTERED"
	CodeDocVerify     Code = "DOCUMENT_VERIFICATION"
	CodeSafetyTrain   Code = "SAFETY_TRAINING"
	CodeMedicalClear  Code = "MEDICAL_CLEARANCE"
	CodeSiteInduction Code = "SITE_INDUCTION"
	CodeApproved      Code = "APPROVED"
	CodeActive        Code = "ACTIVE"
	CodeSuspended     Code = "SUSPENDED"
	CodeExpired       Code = "EXPIRED"
	CodeOffboarded    Code = "OFFBOARDED"
	CodeArchived      Code = "ARCHIVED"
)

// ValidCodes lists all valid state codes.
var ValidCodes = []Code{
	CodeRegistered,
	CodeDocVerify,
	CodeSafetyTrain,
	CodeMedicalClear,
	CodeSiteInduction,
	CodeApproved,
	CodeActive,
	CodeSuspended,
	CodeExpired,
	CodeOffboarded,
	CodeArchived,
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
