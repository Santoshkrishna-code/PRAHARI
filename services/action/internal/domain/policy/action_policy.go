package policy

import (
	"time"

	"prahari/services/action/internal/domain/action"
	"prahari/services/action/internal/domain/escalation"
)

// IsActionOverdue checks if the action's due date has passed.
func IsActionOverdue(act *action.Action) bool {
	if act == nil {
		return false
	}
	return act.Status != "CLOSED" && act.Status != "CANCELLED" && time.Now().After(act.DueDate)
}

// CanCloseAction checks if action is in effectiveness review before closing.
func CanCloseAction(act *action.Action) bool {
	if act == nil {
		return false
	}
	return act.Status == "EFFECTIVENESS_REVIEW"
}

// IsEscalationTriggered checks if overdue days exceed the escalation threshold.
func IsEscalationTriggered(act *action.Action, rule *escalation.Rule) bool {
	if act == nil || rule == nil || !rule.Active {
		return false
	}
	if !IsActionOverdue(act) {
		return false
	}
	overdueDuration := time.Since(act.DueDate)
	overdueDays := int(overdueDuration.Hours() / 24)
	return overdueDays >= rule.OverdueDays
}
