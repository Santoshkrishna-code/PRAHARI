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
			name:    "Planned to Monitoring",
			from:    CodePlanned,
			to:      CodeMonitoring,
			wantErr: false,
		},
		{
			name:    "Monitoring to Sampling",
			from:    CodeMonitoring,
			to:      CodeSampling,
			wantErr: false,
		},
		{
			name:    "Sampling to LaboratoryAnalysis",
			from:    CodeSampling,
			to:      CodeLaboratoryAnalysis,
			wantErr: false,
		},
		{
			name:    "LaboratoryAnalysis to ComplianceEvaluation",
			from:    CodeLaboratoryAnalysis,
			to:      CodeComplianceEvaluation,
			wantErr: false,
		},
		{
			name:    "ComplianceEvaluation to CorrectiveAction",
			from:    CodeComplianceEvaluation,
			to:      CodeCorrectiveAction,
			wantErr: false,
		},
		{
			name:    "CorrectiveAction to Verification",
			from:    CodeCorrectiveAction,
			to:      CodeVerification,
			wantErr: false,
		},
		{
			name:    "Verification to Closed",
			from:    CodeVerification,
			to:      CodeClosed,
			wantErr: false,
		},
		{
			name:    "ComplianceEvaluation to NonCompliant",
			from:    CodeComplianceEvaluation,
			to:      CodeNonCompliant,
			wantErr: false,
		},
		{
			name:    "NonCompliant to Escalated",
			from:    CodeNonCompliant,
			to:      CodeEscalated,
			wantErr: false,
		},
		{
			name:    "Invalid Transition: Planned to LaboratoryAnalysis",
			from:    CodePlanned,
			to:      CodeLaboratoryAnalysis,
			wantErr: true,
		},
		{
			name:    "Unknown State Check",
			from:    Code("UNKNOWN"),
			to:      CodePlanned,
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
