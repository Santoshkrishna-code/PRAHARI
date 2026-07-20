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
			name:    "Planned to Isolation Approved",
			from:    CodePlanned,
			to:      CodeIsolationApproved,
			wantErr: false,
		},
		{
			name:    "Isolation Approved to Locks Applied",
			from:    CodeIsolationApproved,
			to:      CodeLocksApplied,
			wantErr: false,
		},
		{
			name:    "Locks Applied to Tags Applied",
			from:    CodeLocksApplied,
			to:      CodeTagsApplied,
			wantErr: false,
		},
		{
			name:    "Tags Applied to Zero Energy Verified",
			from:    CodeTagsApplied,
			to:      CodeZeroEnergyVerified,
			wantErr: false,
		},
		{
			name:    "Zero Energy Verified to Maintenance In Progress",
			from:    CodeZeroEnergyVerified,
			to:      CodeMaintenanceInProgress,
			wantErr: false,
		},
		{
			name:    "Maintenance In Progress to Restoration Planned",
			from:    CodeMaintenanceInProgress,
			to:      CodeRestorationPlanned,
			wantErr: false,
		},
		{
			name:    "Restoration Planned to Locks Removed",
			from:    CodeRestorationPlanned,
			to:      CodeLocksRemoved,
			wantErr: false,
		},
		{
			name:    "Locks Removed to Returned To Service",
			from:    CodeLocksRemoved,
			to:      CodeReturnedToService,
			wantErr: false,
		},
		{
			name:    "Returned To Service to Closed",
			from:    CodeReturnedToService,
			to:      CodeClosed,
			wantErr: false,
		},
		{
			name:    "Invalid: Planned to Returned To Service",
			from:    CodePlanned,
			to:      CodeReturnedToService,
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
