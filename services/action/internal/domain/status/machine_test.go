package status_test

import (
	"testing"

	"prahari/services/action/internal/domain/status"
)

func TestValidTransitions(t *testing.T) {
	tests := []struct {
		from status.Code
		to   status.Code
	}{
		{status.CodeCreated, status.CodeAssigned},
		{status.CodeCreated, status.CodeCancelled},
		{status.CodeAssigned, status.CodeInProgress},
		{status.CodeAssigned, status.CodeOverdue},
		{status.CodeAssigned, status.CodeCancelled},
		{status.CodeInProgress, status.CodeEvidenceSubmitted},
		{status.CodeInProgress, status.CodeOverdue},
		{status.CodeInProgress, status.CodeCancelled},
		{status.CodeEvidenceSubmitted, status.CodeEffectivenessReview},
		{status.CodeEvidenceSubmitted, status.CodeRejected},
		{status.CodeEffectivenessReview, status.CodeClosed},
		{status.CodeEffectivenessReview, status.CodeRejected},
		{status.CodeOverdue, status.CodeInProgress},
		{status.CodeOverdue, status.CodeCancelled},
		{status.CodeRejected, status.CodeInProgress},
		{status.CodeRejected, status.CodeCancelled},
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
		{status.CodeCreated, status.CodeClosed},
		{status.CodeCreated, status.CodeInProgress},
		{status.CodeAssigned, status.CodeClosed},
		{status.CodeClosed, status.CodeInProgress},
		{status.CodeCancelled, status.CodeAssigned},
		{status.CodeEvidenceSubmitted, status.CodeAssigned},
	}

	for _, tt := range tests {
		t.Run(string(tt.from)+"->"+string(tt.to), func(t *testing.T) {
			if err := status.ValidateTransition(tt.from, tt.to); err == nil {
				t.Errorf("expected transition %s -> %s to be invalid, but got nil", tt.from, tt.to)
			}
		})
	}
}
