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
			name:    "Registered to Scheduled",
			from:    CodeRegistered,
			to:      CodeScheduled,
			wantErr: false,
		},
		{
			name:    "Scheduled to Calibration Started",
			from:    CodeScheduled,
			to:      CodeCalibrationStarted,
			wantErr: false,
		},
		{
			name:    "Calibration Started to Measurement Recorded",
			from:    CodeCalibrationStarted,
			to:      CodeMeasurementRecorded,
			wantErr: false,
		},
		{
			name:    "Measurement Recorded to Tolerance Verification",
			from:    CodeMeasurementRecorded,
			to:      CodeToleranceVerification,
			wantErr: false,
		},
		{
			name:    "Tolerance Verification to Certificate Generated",
			from:    CodeToleranceVerification,
			to:      CodeCertificateGenerated,
			wantErr: false,
		},
		{
			name:    "Certificate Generated to Approved",
			from:    CodeCertificateGenerated,
			to:      CodeApproved,
			wantErr: false,
		},
		{
			name:    "Approved to Active",
			from:    CodeApproved,
			to:      CodeActive,
			wantErr: false,
		},
		{
			name:    "Active to Scheduled",
			from:    CodeActive,
			to:      CodeScheduled,
			wantErr: false,
		},
		{
			name:    "Invalid: Active to Calibration Started",
			from:    CodeActive,
			to:      CodeCalibrationStarted,
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
