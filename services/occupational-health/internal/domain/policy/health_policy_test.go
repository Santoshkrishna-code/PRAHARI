package policy

import (
	"testing"

	"prahari/services/occupational-health/internal/domain/exposure"
	"prahari/services/occupational-health/internal/domain/restriction"
)

func TestEvaluatePermitEligibility(t *testing.T) {
	tests := []struct {
		name         string
		restrictions []restriction.MedicalRestriction
		permitType   string
		wantErr      bool
	}{
		{
			name:         "No Restrictions",
			restrictions: []restriction.MedicalRestriction{},
			permitType:   "WORKING_AT_HEIGHT",
			wantErr:      false,
		},
		{
			name: "Active Height Restriction - Requesting Height Permit",
			restrictions: []restriction.MedicalRestriction{
				{RestrictionCode: "NO_HEIGHT"},
			},
			permitType: "WORKING_AT_HEIGHT",
			wantErr:    true,
		},
		{
			name: "Active Heavy Lift Restriction - Requesting Height Permit",
			restrictions: []restriction.MedicalRestriction{
				{RestrictionCode: "NO_HEAVY_LIFT"},
			},
			permitType: "WORKING_AT_HEIGHT",
			wantErr:    false,
		},
		{
			name: "Active Heavy Lift Restriction - Requesting Lift Permit",
			restrictions: []restriction.MedicalRestriction{
				{RestrictionCode: "NO_HEAVY_LIFT"},
			},
			permitType: "HEAVY_LIFTING",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := EvaluatePermitEligibility(tt.restrictions, tt.permitType)
			if (err != nil) != tt.wantErr {
				t.Errorf("EvaluatePermitEligibility() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEvaluateExposureLimit(t *testing.T) {
	tests := []struct {
		name   string
		record exposure.ExposureRecord
		want   bool
	}{
		{
			name: "Below Limit",
			record: exposure.ExposureRecord{
				ExposureLevel:  50.0,
				LimitThreshold: 85.0,
			},
			want: false,
		},
		{
			name: "Above Limit",
			record: exposure.ExposureRecord{
				ExposureLevel:  90.0,
				LimitThreshold: 85.0,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EvaluateExposureLimit(&tt.record)
			if got != tt.want {
				t.Errorf("EvaluateExposureLimit() got = %v, want %v", got, tt.want)
			}
		})
	}
}
