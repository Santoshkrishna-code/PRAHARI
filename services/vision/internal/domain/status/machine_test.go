package status_test

import (
	"testing"

	"prahari/services/vision/internal/domain/status"
)

func TestValidTransitions(t *testing.T) {
	tests := []struct {
		from status.Code
		to   status.Code
	}{
		{status.CodeOnline, status.CodeOffline},
		{status.CodeOnline, status.CodeUnreachable},
		{status.CodeOffline, status.CodeOnline},
		{status.CodeOffline, status.CodeUnreachable},
		{status.CodeUnreachable, status.CodeOnline},
		{status.CodeUnreachable, status.CodeOffline},
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
		{status.Code("UNKNOWN"), status.CodeOnline},
	}

	for _, tt := range tests {
		t.Run(string(tt.from)+"->"+string(tt.to), func(t *testing.T) {
			if err := status.ValidateTransition(tt.from, tt.to); err == nil {
				t.Errorf("expected transition %s -> %s to be invalid, but got nil", tt.from, tt.to)
			}
		})
	}
}
