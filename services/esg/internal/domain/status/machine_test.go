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
			name:    "ObjectiveDefined to BaselineEstablished",
			from:    CodeObjectiveDefined,
			to:      CodeBaselineEstablished,
			wantErr: false,
		},
		{
			name:    "BaselineEstablished to DataCollection",
			from:    CodeBaselineEstablished,
			to:      CodeDataCollection,
			wantErr: false,
		},
		{
			name:    "DataCollection to Calculation",
			from:    CodeDataCollection,
			to:      CodeCalculation,
			wantErr: false,
		},
		{
			name:    "Calculation to Validation",
			from:    CodeCalculation,
			to:      CodeValidation,
			wantErr: false,
		},
		{
			name:    "Validation to Disclosure",
			from:    CodeValidation,
			to:      CodeDisclosure,
			wantErr: false,
		},
		{
			name:    "Disclosure to ExecutiveReview",
			from:    CodeDisclosure,
			to:      CodeExecutiveReview,
			wantErr: false,
		},
		{
			name:    "ExecutiveReview to Published",
			from:    CodeExecutiveReview,
			to:      CodePublished,
			wantErr: false,
		},
		{
			name:    "ExecutiveReview to Reopened",
			from:    CodeExecutiveReview,
			to:      CodeReopened,
			wantErr: false,
		},
		{
			name:    "Published to Reopened",
			from:    CodePublished,
			to:      CodeReopened,
			wantErr: false,
		},
		{
			name:    "Reopened to DataCollection",
			from:    CodeReopened,
			to:      CodeDataCollection,
			wantErr: false,
		},
		{
			name:    "Invalid: Published to Validation",
			from:    CodePublished,
			to:      CodeValidation,
			wantErr: true,
		},
		{
			name:    "Unknown State",
			from:    Code("UNKNOWN"),
			to:      CodePublished,
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
