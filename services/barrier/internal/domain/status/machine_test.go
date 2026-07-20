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
			name:    "Registered to Assigned",
			from:    CodeRegistered,
			to:      CodeAssigned,
			wantErr: false,
		},
		{
			name:    "Assigned to Operational",
			from:    CodeAssigned,
			to:      CodeOperational,
			wantErr: false,
		},
		{
			name:    "Operational to Proof Test",
			from:    CodeOperational,
			to:      CodeProofTest,
			wantErr: false,
		},
		{
			name:    "Proof Test to Verified",
			from:    CodeProofTest,
			to:      CodeVerified,
			wantErr: false,
		},
		{
			name:    "Verified to Operational",
			from:    CodeVerified,
			to:      CodeOperational,
			wantErr: false,
		},
		{
			name:    "Operational to Bypassed",
			from:    CodeOperational,
			to:      CodeBypassed,
			wantErr: false,
		},
		{
			name:    "Bypassed to Operational",
			from:    CodeBypassed,
			to:      CodeOperational,
			wantErr: false,
		},
		{
			name:    "Operational to Impaired",
			from:    CodeOperational,
			to:      CodeImpaired,
			wantErr: false,
		},
		{
			name:    "Invalid: Registered to Verified",
			from:    CodeRegistered,
			to:      CodeVerified,
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
