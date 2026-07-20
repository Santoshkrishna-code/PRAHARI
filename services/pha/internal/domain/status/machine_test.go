package status

import (
	"testing"
)

func TestValidateTransition(t *testing.T) {
	tests := []struct {
		name    string
		from    Code
		to      Code
		wantErr bool
	}{
		{
			name:    "Draft to Preparation",
			from:    CodeDraft,
			to:      CodePreparation,
			wantErr: false,
		},
		{
			name:    "Preparation to Study",
			from:    CodePreparation,
			to:      CodeStudy,
			wantErr: false,
		},
		{
			name:    "Study to Risk Evaluation",
			from:    CodeStudy,
			to:      CodeRiskEvaluation,
			wantErr: false,
		},
		{
			name:    "Risk Evaluation to Recommendation",
			from:    CodeRiskEvaluation,
			to:      CodeRecommendation,
			wantErr: false,
		},
		{
			name:    "Recommendation to Approval",
			from:    CodeRecommendation,
			to:      CodeApproval,
			wantErr: false,
		},
		{
			name:    "Approval to Implementation Tracking",
			from:    CodeApproval,
			to:      CodeImplementationTracking,
			wantErr: false,
		},
		{
			name:    "Implementation Tracking to Verification",
			from:    CodeImplementationTracking,
			to:      CodeVerification,
			wantErr: false,
		},
		{
			name:    "Verification to Revalidation Scheduled",
			from:    CodeVerification,
			to:      CodeRevalidationScheduled,
			wantErr: false,
		},
		{
			name:    "Revalidation Scheduled to Closed",
			from:    CodeRevalidationScheduled,
			to:      CodeClosed,
			wantErr: false,
		},
		{
			name:    "Invalid: Draft to Verification",
			from:    CodeDraft,
			to:      CodeVerification,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTransition(tt.from, tt.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTransition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
