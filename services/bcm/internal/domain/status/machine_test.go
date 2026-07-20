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
			name:    "Planned to Business Impact Analysis",
			from:    CodePlanned,
			to:      CodeBusinessImpactAnalysis,
			wantErr: false,
		},
		{
			name:    "BIA to Strategy Development",
			from:    CodeBusinessImpactAnalysis,
			to:      CodeStrategyDevelopment,
			wantErr: false,
		},
		{
			name:    "Strategy Development to Plan Development",
			from:    CodeStrategyDevelopment,
			to:      CodePlanDevelopment,
			wantErr: false,
		},
		{
			name:    "Plan Development to Approval",
			from:    CodePlanDevelopment,
			to:      CodeApproval,
			wantErr: false,
		},
		{
			name:    "Approval to Activation",
			from:    CodeApproval,
			to:      CodeActivation,
			wantErr: false,
		},
		{
			name:    "Activation to Recovery",
			from:    CodeActivation,
			to:      CodeRecovery,
			wantErr: false,
		},
		{
			name:    "Recovery to Review",
			from:    CodeRecovery,
			to:      CodeReview,
			wantErr: false,
		},
		{
			name:    "Review to Continuous Improvement",
			from:    CodeReview,
			to:      CodeContinuousImprovement,
			wantErr: false,
		},
		{
			name:    "Continuous Improvement to Approval",
			from:    CodeContinuousImprovement,
			to:      CodeApproval,
			wantErr: false,
		},
		{
			name:    "Invalid: Planned to Recovery",
			from:    CodePlanned,
			to:      CodeRecovery,
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
