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
			name:    "Draft to Planned",
			from:    CodeDraft,
			to:      CodePlanned,
			wantErr: false,
		},
		{
			name:    "Planned to Scheduled",
			from:    CodePlanned,
			to:      CodeScheduled,
			wantErr: false,
		},
		{
			name:    "Draft to Scheduled (Direct bypass is invalid)",
			from:    CodeDraft,
			to:      CodeScheduled,
			wantErr: true,
		},
		{
			name:    "Active to Renewal",
			from:    CodeActive,
			to:      CodeRenewal,
			wantErr: false,
		},
		{
			name:    "Renewal to Active",
			from:    CodeRenewal,
			to:      CodeActive,
			wantErr: false,
		},
		{
			name:    "Unknown State",
			from:    Code("UNKNOWN"),
			to:      CodeActive,
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
