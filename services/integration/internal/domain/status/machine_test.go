package status_test

import (
	"testing"

	"prahari/services/integration/internal/domain/status"
)

func TestValidTransitions(t *testing.T) {
	tests := []struct {
		from status.Code
		to   status.Code
	}{
		{status.CodeDisconnected, status.CodeConnecting},
		{status.CodeConnecting, status.CodeConnected},
		{status.CodeConnecting, status.CodeFailed},
		{status.CodeConnecting, status.CodeDisconnected},
		{status.CodeConnected, status.CodeDisconnected},
		{status.CodeConnected, status.CodeFailed},
		{status.CodeFailed, status.CodeConnecting},
		{status.CodeFailed, status.CodeDisconnected},
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
		{status.CodeDisconnected, status.CodeConnected},
		{status.CodeDisconnected, status.CodeFailed},
		{status.CodeConnected, status.CodeConnecting},
	}

	for _, tt := range tests {
		t.Run(string(tt.from)+"->"+string(tt.to), func(t *testing.T) {
			if err := status.ValidateTransition(tt.from, tt.to); err == nil {
				t.Errorf("expected transition %s -> %s to be invalid, but got nil", tt.from, tt.to)
			}
		})
	}
}
