package policy

import (
	"fmt"
	"time"

	"prahari/services/chemical/internal/domain/chemical"
	"prahari/services/chemical/internal/domain/compatibility"
	"prahari/services/chemical/internal/domain/container"
	"prahari/services/chemical/internal/domain/ghsclassification"
	"prahari/services/chemical/internal/domain/sds"
	"prahari/services/chemical/internal/domain/storagearea"
)

// IsCompatible checks compatibility between two hazard classes using compatibility rules.
func IsCompatible(classA, classB string, rules []*compatibility.Rule) bool {
	for _, rule := range rules {
		if (rule.ClassA == classA && rule.ClassB == classB) || (rule.ClassA == classB && rule.ClassB == classA) {
			return rule.Compatible
		}
	}
	return true
}

// VerifyMaxAllowableQuantity checks if adding a container exceeds the Storage Area's MAQ.
func VerifyMaxAllowableQuantity(area *storagearea.Area, additionalQty float64) error {
	if area == nil {
		return nil
	}
	if area.CurrentLoadKg+additionalQty > area.MaxCapacityKg {
		return fmt.Errorf("adding %.2f kg exceeds the Maximum Allowable Quantity (MAQ) of %.2f kg for storage area %s",
			additionalQty, area.MaxCapacityKg, area.Name)
	}
	return nil
}

// IsSDSRequiredBeforeIssue checks if the chemical has a valid and active SDS in the repository.
func IsSDSRequiredBeforeIssue(c *chemical.Chemical, activeSds *sds.SDS) bool {
	if c == nil {
		return false
	}
	return activeSds != nil && activeSds.ChemicalID == c.ID && activeSds.DocumentURL != ""
}

// IsGHSClassificationRequired checks if GHS classification is provided.
func IsGHSClassificationRequired(c *chemical.Chemical, g *ghsclassification.GHS) bool {
	if c == nil {
		return false
	}
	return g != nil && g.ChemicalID == c.ID && g.SignalWord != ""
}

// IsExpired checks if a container has passed its expiration date.
func IsExpired(con *container.Container) bool {
	if con == nil || con.ExpiryDate == nil {
		return false
	}
	return time.Now().After(*con.ExpiryDate)
}

// ExceedsOSHAPSMThreshold checks if quantity exceeds OSHA Process Safety Management threshold limits.
func ExceedsOSHAPSMThreshold(c *chemical.Chemical, qty float64) bool {
	if c == nil {
		return false
	}
	// Simplified threshold checking: e.g. for restricted chemicals, check if total qty > 4500 kg
	if c.IsRestricted && qty >= 4500 {
		return true
	}
	return qty >= 10000 // default general hazard threshold
}

// ExceedsSEVESOThreshold checks if volume exceeds European SEVESO III safety directive thresholds.
func ExceedsSEVESOThreshold(c *chemical.Chemical, qty float64) bool {
	if c == nil {
		return false
	}
	if c.IsRestricted && qty >= 5000 {
		return true
	}
	return qty >= 15000
}
