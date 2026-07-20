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
			name:    "Draft to Review",
			from:    CodeDraft,
			to:      CodeReview,
			wantErr: false,
		},
		{
			name:    "Review to Approval",
			from:    CodeReview,
			to:      CodeApproval,
			wantErr: false,
		},
		{
			name:    "Approval to Published",
			from:    CodeApproval,
			to:      CodePublished,
			wantErr: false,
		},
		{
			name:    "Published to Controlled Distribution",
			from:    CodePublished,
			to:      CodeControlledDistribution,
			wantErr: false,
		},
		{
			name:    "Controlled Distribution to Periodic Review",
			from:    CodeControlledDistribution,
			to:      CodePeriodicReview,
			wantErr: false,
		},
		{
			name:    "Periodic Review to Revision",
			from:    CodePeriodicReview,
			to:      CodeRevision,
			wantErr: false,
		},
		{
			name:    "Revision to Draft",
			from:    CodeRevision,
			to:      CodeDraft,
			wantErr: false,
		},
		{
			name:    "Published to Superseded",
			from:    CodePublished,
			to:      CodeSuperseded,
			wantErr: false,
		},
		{
			name:    "Superseded to Archived",
			from:    CodeSuperseded,
			to:      CodeArchived,
			wantErr: false,
		},
		{
			name:    "Invalid: Draft to Archived",
			from:    CodeDraft,
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
