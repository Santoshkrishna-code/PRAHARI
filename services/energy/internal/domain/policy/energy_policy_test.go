package policy

import (
	"testing"

	"prahari/services/energy/internal/domain/energyconsumption"
)

func TestEvaluateEnergyTarget(t *testing.T) {
	tests := []struct {
		name        string
		consumption energyconsumption.Consumption
		limit       float64
		want        bool
	}{
		{
			name: "Within Target",
			consumption: energyconsumption.Consumption{
				ConsumptionKWh: 45000.0,
			},
			limit: 50000.0,
			want:  true,
		},
		{
			name: "Exceeded Target",
			consumption: energyconsumption.Consumption{
				ConsumptionKWh: 55000.0,
			},
			limit: 50000.0,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EvaluateEnergyTarget(&tt.consumption, tt.limit)
			if got != tt.want {
				t.Errorf("EvaluateEnergyTarget() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertConsumptionToCarbon(t *testing.T) {
	t.Run("Standard Grid Factor", func(t *testing.T) {
		got := ConvertConsumptionToCarbon(1000.0, 0.45)
		want := 450.0
		if got != want {
			t.Errorf("ConvertConsumptionToCarbon() got = %v, want %v", got, want)
		}
	})

	t.Run("Zero Consumption", func(t *testing.T) {
		got := ConvertConsumptionToCarbon(0.0, 0.45)
		want := 0.0
		if got != want {
			t.Errorf("ConvertConsumptionToCarbon() got = %v, want %v", got, want)
		}
	})
}
