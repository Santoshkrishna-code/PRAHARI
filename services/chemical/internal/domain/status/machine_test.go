package status_test

import (
	"testing"

	"prahari/services/chemical/internal/domain/status"
)

func TestValidTransitions(t *testing.T) {
	tests := []struct {
		from status.Code
		to   status.Code
	}{
		{status.CodeRequested, status.CodeTechnicalReview},
		{status.CodeRequested, status.CodeRejected},
		{status.CodeTechnicalReview, status.CodeSafetyReview},
		{status.CodeTechnicalReview, status.CodeRejected},
		{status.CodeSafetyReview, status.CodeEnvironmentalReview},
		{status.CodeSafetyReview, status.CodeRejected},
		{status.CodeEnvironmentalReview, status.CodeApproved},
		{status.CodeEnvironmentalReview, status.CodeRejected},
		{status.CodeApproved, status.CodePurchased},
		{status.CodeApproved, status.CodeRecalled},
		{status.CodePurchased, status.CodeReceived},
		{status.CodeReceived, status.CodeInspection},
		{status.CodeInspection, status.CodeStored},
		{status.CodeStored, status.CodeIssued},
		{status.CodeStored, status.CodeTransferred},
		{status.CodeIssued, status.CodeInUse},
		{status.CodeIssued, status.CodeReturned},
		{status.CodeInUse, status.CodeWasteClassification},
		{status.CodeTransferred, status.CodeStored},
		{status.CodeReturned, status.CodeStored},
		{status.CodeWasteClassification, status.CodeWasteProcessing},
		{status.CodeWasteProcessing, status.CodeDisposed},
		{status.CodeDisposed, status.CodeClosed},
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
		{status.CodeRequested, status.CodeApproved},
		{status.CodeApproved, status.CodeClosed},
		{status.CodeDisposed, status.CodeStored},
		{status.CodeClosed, status.CodeRequested},
		{status.CodeRejected, status.CodeRequested},
	}

	for _, tt := range tests {
		t.Run(string(tt.from)+"->"+string(tt.to), func(t *testing.T) {
			if err := status.ValidateTransition(tt.from, tt.to); err == nil {
				t.Errorf("expected transition %s -> %s to be invalid, but got nil", tt.from, tt.to)
			}
		})
	}
}
