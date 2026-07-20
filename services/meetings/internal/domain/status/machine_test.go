package status_test

import (
	"testing"

	"prahari/services/meetings/internal/domain/status"
)

func TestValidTransitions(t *testing.T) {
	tests := []struct {
		from status.Code
		to   status.Code
	}{
		{status.CodePlanned, status.CodeScheduled},
		{status.CodePlanned, status.CodeCancelled},
		{status.CodeScheduled, status.CodeInProgress},
		{status.CodeScheduled, status.CodeRescheduled},
		{status.CodeScheduled, status.CodeCancelled},
		{status.CodeInProgress, status.CodeAttendanceRecorded},
		{status.CodeInProgress, status.CodeCancelled},
		{status.CodeAttendanceRecorded, status.CodeMinutesApproved},
		{status.CodeAttendanceRecorded, status.CodeCancelled},
		{status.CodeMinutesApproved, status.CodeActionsGenerated},
		{status.CodeMinutesApproved, status.CodeClosed},
		{status.CodeActionsGenerated, status.CodeClosed},
		{status.CodeRescheduled, status.CodeScheduled},
		{status.CodeRescheduled, status.CodeCancelled},
	}

	for _, tt := range tests {
		t.Run(string(tt.from)+"->"+string(tt.to), func(t *testing.T) {
			if err := status.ValidateTransition(tt.from, tt.to); err != nil {
				t.Errorf("expected transition %s -> %s to be valid, got error: %v", tt.from, tt.to, err)
			}
		})
	}
}

func TestInvalidTransitions(t *testing.T) {
	tests := []struct {
		from status.Code
		to   status.Code
	}{
		{status.CodePlanned, status.CodeClosed},
		{status.CodePlanned, status.CodeInProgress},
		{status.CodeScheduled, status.CodeClosed},
		{status.CodeClosed, status.CodeInProgress},
		{status.CodeCancelled, status.CodeScheduled},
		{status.CodeInProgress, status.CodeClosed},
		{status.CodeActionsGenerated, status.CodeInProgress},
	}

	for _, tt := range tests {
		t.Run(string(tt.from)+"->"+string(tt.to), func(t *testing.T) {
			if err := status.ValidateTransition(tt.from, tt.to); err == nil {
				t.Errorf("expected transition %s -> %s to be invalid, but got nil", tt.from, tt.to)
			}
		})
	}
}
