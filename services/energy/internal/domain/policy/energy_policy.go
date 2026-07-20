package policy

import (
	"prahari/services/energy/internal/domain/energyconsumption"
)

// EvaluateEnergyTarget checks if facility consumption exceeded target limit.
func EvaluateEnergyTarget(c *energyconsumption.Consumption, limitKWh float64) bool {
	return c.ConsumptionKWh <= limitKWh
}

// ConvertConsumptionToCarbon calculates Scope 2 equivalent emissions.
func ConvertConsumptionToCarbon(kwh float64, gridFactor float64) float64 {
	return kwh * gridFactor
}
