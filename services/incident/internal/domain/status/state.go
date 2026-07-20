package status

import (
	"fmt"
)

// Code represents the lifecycle state of an incident.
type Code string

const (
	CodeDraft         Code = "DRAFT"
	CodeSubmitted     Code = "SUBMITTED"
	CodeUnderReview   Code = "UNDER_REVIEW"
	CodeAssigned      Code = "ASSIGNED"
	CodeInvestigating Code = "INVESTIGATING"
	CodeCAPAInProgress Code = "CAPA_IN_PROGRESS"
	CodeResolved      Code = "RESOLVED"
	CodeClosed        Code = "CLOSED"
	CodeRejected      Code = "REJECTED"
	CodeCancelled     Code = "CANCELLED"
)

// ValidCodes enumerates all accepted lifecycle states.
var ValidCodes = []Code{
	CodeDraft,
	CodeSubmitted,
	CodeUnderReview,
	CodeAssigned,
	CodeInvestigating,
	CodeCAPAInProgress,
	CodeResolved,
	CodeClosed,
	CodeRejected,
	CodeCancelled,
}

// Validate checks whether the code is among accepted lifecycle states.
func (c Code) Validate() error {
	for _, valid := range ValidCodes {
		if c == valid {
			return nil
		}
	}
	return fmt.Errorf("invalid status code: %s", c)
}

// IsTerminal returns true if the status represents a final lifecycle state.
func (c Code) IsTerminal() bool {
	return c == CodeClosed || c == CodeRejected || c == CodeCancelled
}

// IsActive returns true if the incident is in an active, non-terminal state.
func (c Code) IsActive() bool {
	return !c.IsTerminal()
}

// RequiresAssignee returns true if the status requires an assigned user.
func (c Code) RequiresAssignee() bool {
	return c == CodeAssigned || c == CodeInvestigating || c == CodeCAPAInProgress
}
