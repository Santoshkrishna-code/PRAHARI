package policy

import (
	"prahari/services/water/internal/domain/consumption"
)

// EvaluateWaterTarget evaluates whether water consumption exceeds budget threshold.
func EvaluateWaterTarget(c *consumption.Consumption, budgetKL float64) bool {
	if budgetKL <= 0 {
		return true
	}
	return c.ConsumptionKL <= budgetKL
}

// CalculateWaterBalance computes loss percent between total intake and sum of metered consumption.
func CalculateWaterBalance(totalIntakeKL, totalMeteredKL float64) (lossKL float64, lossPct float64) {
	if totalIntakeKL <= 0 {
		return 0, 0
	}
	lossKL = totalIntakeKL - totalMeteredKL
	if lossKL < 0 {
		lossKL = 0
	}
	lossPct = (lossKL / totalIntakeKL) * 100.0
	return lossKL, lossPct
}

// CalculateRecycleRatio returns recycled water volume percentage of total water used.
func CalculateRecycleRatio(recycledKL, totalUsageKL float64) float64 {
	if totalUsageKL <= 0 {
		return 0
	}
	return (recycledKL / totalUsageKL) * 100.0
}
