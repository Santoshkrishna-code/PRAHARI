package policy

import (
	"testing"

	"prahari/services/esg/internal/domain/carboninventory"
)

func TestEvaluateCarbonTargets(t *testing.T) {
	tests := []struct {
		name      string
		inventory carboninventory.Inventory
		target    float64
		want      bool
	}{
		{
			name: "Goal Achieved (Within Target)",
			inventory: carboninventory.Inventory{
				TotalCo2Kg: 45000.0,
			},
			target: 50000.0,
			want:   true,
		},
		{
			name: "Goal Failed (Above Target)",
			inventory: carboninventory.Inventory{
				TotalCo2Kg: 55000.0,
			},
			target: 50000.0,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EvaluateCarbonTargets(&tt.inventory, tt.target)
			if got != tt.want {
				t.Errorf("EvaluateCarbonTargets() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConversions(t *testing.T) {
	t.Run("ConvertElectricityToScope2", func(t *testing.T) {
		got := ConvertElectricityToScope2(1000.0, 0.45)
		want := 450.0
		if got != want {
			t.Errorf("ConvertElectricityToScope2() got = %v, want %v", got, want)
		}
	})

	t.Run("ConvertFuelToScope1", func(t *testing.T) {
		got := ConvertFuelToScope1(500.0, 2.68)
		want := 1340.0
		if got != want {
			t.Errorf("ConvertFuelToScope1() got = %v, want %v", got, want)
		}
	})
}
