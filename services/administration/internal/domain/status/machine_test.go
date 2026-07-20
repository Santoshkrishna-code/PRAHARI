package status_test

import (
	"testing"

	"prahari/services/administration/internal/domain/status"
)

func TestValidTransitions(t *testing.T) {
	tests := []struct {
		from status.Code
		to   status.Code
	}{
		{status.CodeDraft, status.CodeConfigured},
		{status.CodeDraft, status.CodeArchived},
		{status.CodeConfigured, status.CodeValidated},
		{status.CodeConfigured, status.CodeArchived},
		{status.CodeValidated, status.CodeActivated},
		{status.CodeValidated, status.CodeArchived},
		{status.CodeActivated, status.CodeOperational},
		{status.CodeActivated, status.CodeSuspended},
		{status.CodeActivated, status.CodeArchived},
		{status.CodeOperational, status.CodeSuspended},
		{status.CodeOperational, status.CodeArchived},
		{status.CodeSuspended, status.CodeOperational},
		{status.CodeSuspended, status.CodeArchived},
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
		{status.CodeDraft, status.CodeOperational},
		{status.CodeOperational, status.CodeValidated},
		{status.CodeArchived, status.CodeDraft},
	}

	for _, tt := range tests {
		t.Run(string(tt.from)+"->"+string(tt.to), func(t *testing.T) {
			if err := status.ValidateTransition(tt.from, tt.to); err == nil {
				t.Errorf("expected transition %s -> %s to be invalid, but got nil", tt.from, tt.to)
			}
		})
	}
}
