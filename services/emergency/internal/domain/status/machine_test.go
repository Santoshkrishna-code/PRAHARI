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
			name:    "Prepared to Declared",
			from:    CodePrepared,
			to:      CodeDeclared,
			wantErr: false,
		},
		{
			name:    "Declared to Response Activated",
			from:    CodeDeclared,
			to:      CodeResponseActivated,
			wantErr: false,
		},
		{
			name:    "Response Activated to Command Established",
			from:    CodeResponseActivated,
			to:      CodeCommandEstablished,
			wantErr: false,
		},
		{
			name:    "Command Established to Resource Deployment",
			from:    CodeCommandEstablished,
			to:      CodeResourceDeployment,
			wantErr: false,
		},
		{
			name:    "Resource Deployment to Evacuation",
			from:    CodeResourceDeployment,
			to:      CodeEvacuation,
			wantErr: false,
		},
		{
			name:    "Evacuation to Stabilized",
			from:    CodeEvacuation,
			to:      CodeStabilized,
			wantErr: false,
		},
		{
			name:    "Stabilized to Recovery",
			from:    CodeStabilized,
			to:      CodeRecovery,
			wantErr: false,
		},
		{
			name:    "Recovery to After Action Review",
			from:    CodeRecovery,
			to:      CodeAfterActionReview,
			wantErr: false,
		},
		{
			name:    "After Action Review to Closed",
			from:    CodeAfterActionReview,
			to:      CodeClosed,
			wantErr: false,
		},
		{
			name:    "Invalid: Prepared to Recovery",
			from:    CodePrepared,
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
