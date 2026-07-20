package policy

import (
	"prahari/services/esg/internal/domain/carboninventory"
)

// EvaluateCarbonTargets evaluates progress of carbon inventory runs.
func EvaluateCarbonTargets(inv *carboninventory.Inventory, target float64) bool {
	return inv.TotalCo2Kg <= target
}

// ConvertElectricityToScope2 co2 calculates electricity scope.
func ConvertElectricityToScope2(kwh float64, gridFactor float64) float64 {
	return kwh * gridFactor
}

// ConvertFuelToScope1 co2 calculates fuel direct scope.
func ConvertFuelToScope1(liters float64, fuelFactor float64) float64 {
	return liters * fuelFactor
}
