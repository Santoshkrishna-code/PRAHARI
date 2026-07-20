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
			name:    "SourceRegistered to Collection",
			from:    CodeSourceRegistered,
			to:      CodeCollection,
			wantErr: false,
		},
		{
			name:    "Collection to Storage",
			from:    CodeCollection,
			to:      CodeStorage,
			wantErr: false,
		},
		{
			name:    "Storage to Treatment",
			from:    CodeStorage,
			to:      CodeTreatment,
			wantErr: false,
		},
		{
			name:    "Treatment to Distribution",
			from:    CodeTreatment,
			to:      CodeDistribution,
			wantErr: false,
		},
		{
			name:    "Distribution to Consumption",
			from:    CodeDistribution,
			to:      CodeConsumption,
			wantErr: false,
		},
		{
			name:    "Consumption to Recycling/Reuse",
			from:    CodeConsumption,
			to:      CodeRecyclingReuse,
			wantErr: false,
		},
		{
			name:    "Recycling/Reuse to Performance Review",
			from:    CodeRecyclingReuse,
			to:      CodePerformanceReview,
			wantErr: false,
		},
		{
			name:    "Performance Review to Archived",
			from:    CodePerformanceReview,
			to:      CodeArchived,
			wantErr: false,
		},
		{
			name:    "Invalid: SourceRegistered to Archived",
			from:    CodeSourceRegistered,
			to:      CodeArchived,
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
