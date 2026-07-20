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
			name:    "Scheduled to Host Approval",
			from:    CodeScheduled,
			to:      CodeHostApproval,
			wantErr: false,
		},
		{
			name:    "Host Approval to Security Verification",
			from:    CodeHostApproval,
			to:      CodeSecurityVerification,
			wantErr: false,
		},
		{
			name:    "Security Verification to Gate Pass Issued",
			from:    CodeSecurityVerification,
			to:      CodeGatePassIssued,
			wantErr: false,
		},
		{
			name:    "Gate Pass Issued to Checked In",
			from:    CodeGatePassIssued,
			to:      CodeCheckedIn,
			wantErr: false,
		},
		{
			name:    "Checked In to On Site",
			from:    CodeCheckedIn,
			to:      CodeOnSite,
			wantErr: false,
		},
		{
			name:    "On Site to Checked Out",
			from:    CodeOnSite,
			to:      CodeCheckedOut,
			wantErr: false,
		},
		{
			name:    "Checked Out to Closed",
			from:    CodeCheckedOut,
			to:      CodeClosed,
			wantErr: false,
		},
		{
			name:    "Scheduled to Cancelled",
			from:    CodeScheduled,
			to:      CodeCancelled,
			wantErr: false,
		},
		{
			name:    "Scheduled to Blacklisted",
			from:    CodeScheduled,
			to:      CodeBlacklisted,
			wantErr: false,
		},
		{
			name:    "Invalid: Scheduled to Closed",
			from:    CodeScheduled,
			to:      CodeClosed,
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
