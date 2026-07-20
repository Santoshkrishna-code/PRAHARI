package policy_test

import (
	"testing"
	"time"

	"prahari/services/action/internal/domain/action"
	"prahari/services/action/internal/domain/escalation"
	"prahari/services/action/internal/domain/policy"
)

func TestIsActionOverdue(t *testing.T) {
	overdueAction := &action.Action{
		Status:  "IN_PROGRESS",
		DueDate: time.Now().Add(-48 * time.Hour),
	}
	if !policy.IsActionOverdue(overdueAction) {
		t.Error("expected action to be overdue")
	}

	futureAction := &action.Action{
		Status:  "IN_PROGRESS",
		DueDate: time.Now().Add(48 * time.Hour),
	}
	if policy.IsActionOverdue(futureAction) {
		t.Error("expected action to not be overdue")
	}

	closedOverdue := &action.Action{
		Status:  "CLOSED",
		DueDate: time.Now().Add(-48 * time.Hour),
	}
	if policy.IsActionOverdue(closedOverdue) {
		t.Error("expected closed action to not be considered overdue")
	}
}

func TestCanCloseAction(t *testing.T) {
	reviewAction := &action.Action{Status: "EFFECTIVENESS_REVIEW"}
	if !policy.CanCloseAction(reviewAction) {
		t.Error("expected action in EFFECTIVENESS_REVIEW to be closable")
	}

	progressAction := &action.Action{Status: "IN_PROGRESS"}
	if policy.CanCloseAction(progressAction) {
		t.Error("expected action in IN_PROGRESS to not be closable")
	}
}

func TestIsEscalationTriggered(t *testing.T) {
	overdueAct := &action.Action{
		Status:  "IN_PROGRESS",
		DueDate: time.Now().Add(-10 * 24 * time.Hour),
	}

	rule := &escalation.Rule{
		OverdueDays: 7,
		Active:      true,
	}

	if !policy.IsEscalationTriggered(overdueAct, rule) {
		t.Error("expected escalation to be triggered (10 overdue days >= 7 threshold)")
	}

	shortRule := &escalation.Rule{
		OverdueDays: 14,
		Active:      true,
	}
	if policy.IsEscalationTriggered(overdueAct, shortRule) {
		t.Error("expected escalation to not trigger (10 overdue days < 14 threshold)")
	}

	inactiveRule := &escalation.Rule{
		OverdueDays: 7,
		Active:      false,
	}
	if policy.IsEscalationTriggered(overdueAct, inactiveRule) {
		t.Error("expected escalation to not trigger on inactive rule")
	}
}
