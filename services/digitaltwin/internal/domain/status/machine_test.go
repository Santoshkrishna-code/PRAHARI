package status_test

import (
	"testing"

	"prahari/services/digitaltwin/internal/domain/status"
)

func TestValidTransitions(t *testing.T) {
	tests := []struct {
		from status.Code
		to   status.Code
	}{
		{status.CodeDraft, status.CodeActive},
		{status.CodeActive, status.CodeArchived},
		{status.CodeArchived, status.CodeActive},
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
		{status.CodeDraft, status.CodeArchived},
		{status.CodeArchived, status.CodeDraft},
		{status.Code("UNKNOWN"), status.CodeActive},
	}

	for _, tt := range tests {
		t.Run(string(tt.from)+"->"+string(tt.to), func(t *testing.T) {
			if err := status.ValidateTransition(tt.from, tt.to); err == nil {
				t.Errorf("expected transition %s -> %s to be invalid, but got nil", tt.from, tt.to)
			}
		})
	}
}
