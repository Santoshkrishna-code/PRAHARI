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
			name:    "MeterRegistered to DataCollection",
			from:    CodeMeterRegistered,
			to:      CodeDataCollection,
			wantErr: false,
		},
		{
			name:    "DataCollection to Validation",
			from:    CodeDataCollection,
			to:      CodeValidation,
			wantErr: false,
		},
		{
			name:    "Validation to Aggregation",
			from:    CodeValidation,
			to:      CodeAggregation,
			wantErr: false,
		},
		{
			name:    "Aggregation to Analysis",
			from:    CodeAggregation,
			to:      CodeAnalysis,
			wantErr: false,
		},
		{
			name:    "Analysis to Optimization",
			from:    CodeAnalysis,
			to:      CodeOptimization,
			wantErr: false,
		},
		{
			name:    "Analysis to Reporting (skip optimization)",
			from:    CodeAnalysis,
			to:      CodeReporting,
			wantErr: false,
		},
		{
			name:    "Optimization to Reporting",
			from:    CodeOptimization,
			to:      CodeReporting,
			wantErr: false,
		},
		{
			name:    "Reporting to Archived",
			from:    CodeReporting,
			to:      CodeArchived,
			wantErr: false,
		},
		{
			name:    "Invalid: MeterRegistered to Reporting",
			from:    CodeMeterRegistered,
			to:      CodeReporting,
			wantErr: true,
		},
		{
			name:    "Unknown State",
			from:    Code("UNKNOWN"),
			to:      CodeDataCollection,
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
