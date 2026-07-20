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
			name:    "Draft to Impact Assessment",
			from:    CodeDraft,
			to:      CodeImpactAssessment,
			wantErr: false,
		},
		{
			name:    "Impact Assessment to Technical Review",
			from:    CodeImpactAssessment,
			to:      CodeTechnicalReview,
			wantErr: false,
		},
		{
			name:    "Technical Review to Risk Review",
			from:    CodeTechnicalReview,
			to:      CodeRiskReview,
			wantErr: false,
		},
		{
			name:    "Risk Review to Safety Review",
			from:    CodeRiskReview,
			to:      CodeSafetyReview,
			wantErr: false,
		},
		{
			name:    "Safety Review to Approval",
			from:    CodeSafetyReview,
			to:      CodeApproval,
			wantErr: false,
		},
		{
			name:    "Approval to Implementation",
			from:    CodeApproval,
			to:      CodeImplementation,
			wantErr: false,
		},
		{
			name:    "Implementation to Verification",
			from:    CodeImplementation,
			to:      CodeVerification,
			wantErr: false,
		},
		{
			name:    "Verification to Closeout",
			from:    CodeVerification,
			to:      CodeCloseout,
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
