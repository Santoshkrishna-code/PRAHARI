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
			name:    "Scheduled to Crew Assigned",
			from:    CodeScheduled,
			to:      CodeCrewAssigned,
			wantErr: false,
		},
		{
			name:    "Crew Assigned to Shift Started",
			from:    CodeCrewAssigned,
			to:      CodeShiftStarted,
			wantErr: false,
		},
		{
			name:    "Shift Started to Operational",
			from:    CodeShiftStarted,
			to:      CodeOperational,
			wantErr: false,
		},
		{
			name:    "Operational to Handover Initiated",
			from:    CodeOperational,
			to:      CodeHandoverInitiated,
			wantErr: false,
		},
		{
			name:    "Handover Initiated to Handover Accepted",
			from:    CodeHandoverInitiated,
			to:      CodeHandoverAccepted,
			wantErr: false,
		},
		{
			name:    "Handover Accepted to Shift Closed",
			from:    CodeHandoverAccepted,
			to:      CodeShiftClosed,
			wantErr: false,
		},
		{
			name:    "Scheduled to Cancelled",
			from:    CodeScheduled,
			to:      CodeCancelled,
			wantErr: false,
		},
		{
			name:    "Invalid: Scheduled to Operational",
			from:    CodeScheduled,
			to:      CodeOperational,
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
