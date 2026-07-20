package policy

import (
	"testing"

	"prahari/services/water/internal/domain/consumption"
)

func TestEvaluateWaterTarget(t *testing.T) {
	c := &consumption.Consumption{ConsumptionKL: 5000.0}
	if !EvaluateWaterTarget(c, 10000.0) {
		t.Errorf("Expected consumption to be within budget")
	}
	if EvaluateWaterTarget(c, 3000.0) {
		t.Errorf("Expected consumption to exceed budget")
	}
}

func TestCalculateWaterBalance(t *testing.T) {
	lossKL, lossPct := CalculateWaterBalance(10000.0, 9500.0)
	if lossKL != 500.0 || lossPct != 5.0 {
		t.Errorf("CalculateWaterBalance() got lossKL=%.2f, lossPct=%.2f; want 500.0, 5.0", lossKL, lossPct)
	}
}

func TestCalculateRecycleRatio(t *testing.T) {
	ratio := CalculateRecycleRatio(4000.0, 10000.0)
	if ratio != 40.0 {
		t.Errorf("CalculateRecycleRatio() got %.2f, want 40.0", ratio)
	}
}
