package policy

import (
	"testing"

	"prahari/services/environmental/internal/domain/emission"
	"prahari/services/environmental/internal/domain/waterquality"
)

func TestEvaluateEmissionLimits(t *testing.T) {
	tests := []struct {
		name    string
		reading emission.Emission
		wantErr bool
	}{
		{
			name: "Normal Emission Level",
			reading: emission.Emission{
				ReleaseRate:    15.5,
				LimitThreshold: 50.0,
			},
			wantErr: false,
		},
		{
			name: "Exceeded Emission Level",
			reading: emission.Emission{
				ReleaseRate:    55.2,
				LimitThreshold: 50.0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := EvaluateEmissionLimits(&tt.reading)
			if (err != nil) != tt.wantErr {
				t.Errorf("EvaluateEmissionLimits() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEvaluateWaterQuality(t *testing.T) {
	tests := []struct {
		name    string
		reading waterquality.WaterQuality
		wantErr bool
	}{
		{
			name: "Normal Water parameters",
			reading: waterquality.WaterQuality{
				PH:           7.2,
				TurbidityNTU: 1.5,
			},
			wantErr: false,
		},
		{
			name: "Acidic pH levels",
			reading: waterquality.WaterQuality{
				PH:           5.2,
				TurbidityNTU: 1.5,
			},
			wantErr: true,
		},
		{
			name: "Alkaline pH levels",
			reading: waterquality.WaterQuality{
				PH:           9.5,
				TurbidityNTU: 1.5,
			},
			wantErr: true,
		},
		{
			name: "High Turbidity levels",
			reading: waterquality.WaterQuality{
				PH:           7.2,
				TurbidityNTU: 8.5,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := EvaluateWaterQuality(&tt.reading)
			if (err != nil) != tt.wantErr {
				t.Errorf("EvaluateWaterQuality() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
