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
			name:    "Available to Reserved",
			from:    CodeAvailable,
			to:      CodeReserved,
			wantErr: false,
		},
		{
			name:    "Reserved to Issued",
			from:    CodeReserved,
			to:      CodeIssued,
			wantErr: false,
		},
		{
			name:    "Issued to In Use",
			from:    CodeIssued,
			to:      CodeInUse,
			wantErr: false,
		},
		{
			name:    "In Use to Inspection",
			from:    CodeInUse,
			to:      CodeInspection,
			wantErr: false,
		},
		{
			name:    "Inspection to Returned",
			from:    CodeInspection,
			to:      CodeReturned,
			wantErr: false,
		},
		{
			name:    "Returned to Available",
			from:    CodeReturned,
			to:      CodeAvailable,
			wantErr: false,
		},
		{
			name:    "Inspection to Maintenance",
			from:    CodeInspection,
			to:      CodeMaintenance,
			wantErr: false,
		},
		{
			name:    "Maintenance to Available",
			from:    CodeMaintenance,
			to:      CodeAvailable,
			wantErr: false,
		},
		{
			name:    "Available to Expired",
			from:    CodeAvailable,
			to:      CodeExpired,
			wantErr: false,
		},
		{
			name:    "Expired to Disposed",
			from:    CodeExpired,
			to:      CodeDisposed,
			wantErr: false,
		},
		{
			name:    "Invalid: Available to Returned",
			from:    CodeAvailable,
			to:      CodeReturned,
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
