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
			name:    "Scheduled to MedicalExamination",
			from:    CodeScheduled,
			to:      CodeMedicalExamination,
			wantErr: false,
		},
		{
			name:    "MedicalExamination to LaboratoryTesting",
			from:    CodeMedicalExamination,
			to:      CodeLaboratoryTesting,
			wantErr: false,
		},
		{
			name:    "Scheduled to ActiveMonitoring (Direct bypass is invalid)",
			from:    CodeScheduled,
			to:      CodeActiveMonitoring,
			wantErr: true,
		},
		{
			name:    "MedicalClearance to ActiveMonitoring",
			from:    CodeMedicalClearance,
			to:      CodeActiveMonitoring,
			wantErr: false,
		},
		{
			name:    "ActiveMonitoring to Restricted",
			from:    CodeActiveMonitoring,
			to:      CodeRestricted,
			wantErr: false,
		},
		{
			name:    "TemporarilyUnfit to Scheduled",
			from:    CodeTemporarilyUnfit,
			to:      CodeScheduled,
			wantErr: false,
		},
		{
			name:    "PermanentlyUnfit to Scheduled (Invalid transition)",
			from:    CodePermanentlyUnfit,
			to:      CodeScheduled,
			wantErr: true,
		},
		{
			name:    "Unknown State",
			from:    Code("UNKNOWN"),
			to:      CodeScheduled,
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
