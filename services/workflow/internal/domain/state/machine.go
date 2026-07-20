package state

import (
	"fmt"
)

// Allowed states values
const (
	StateCreated      = "CREATED"
	StateRunning      = "RUNNING"
	StateWaiting      = "WAITING"
	StateApproved     = "APPROVED"
	StateCompleted    = "COMPLETED"
	StateFailed       = "FAILED"
	StateCompensating = "COMPENSATING"
)

// VerifyTransition validates explicit transition rules.
func VerifyTransition(from, to string) error {
	switch from {
	case StateCreated:
		if to == StateRunning || to == StateFailed {
			return nil
		}
	case StateRunning:
		if to == StateWaiting || to == StateCompleted || to == StateFailed || to == StateCompensating {
			return nil
		}
	case StateWaiting:
		if to == StateApproved || to == StateRunning || to == StateFailed {
			return nil
		}
	case StateApproved:
		if to == StateRunning || to == StateCompleted || to == StateFailed {
			return nil
		}
	}

	return fmt.Errorf("state machine: invalid transition from %s to %s", from, to)
}
