package policy_test

import (
	"testing"

	"prahari/services/chemical/internal/domain/chemical"
	"prahari/services/chemical/internal/domain/compatibility"
	"prahari/services/chemical/internal/domain/policy"
	"prahari/services/chemical/internal/domain/storagearea"
)

func TestIsCompatible(t *testing.T) {
	rules := []*compatibility.Rule{
		{ClassA: "ACID", ClassB: "BASE", Compatible: false},
		{ClassA: "FLAMMABLE", ClassB: "OXIDISER", Compatible: false},
	}

	if policy.IsCompatible("ACID", "BASE", rules) {
		t.Error("expected ACID and BASE to be incompatible")
	}

	if !policy.IsCompatible("ACID", "FLAMMABLE", rules) {
		t.Error("expected ACID and FLAMMABLE to be compatible by default")
	}
}

func TestVerifyMaxAllowableQuantity(t *testing.T) {
	area := &storagearea.Area{
		Name:          "Solvent Locker",
		MaxCapacityKg: 100.0,
		CurrentLoadKg: 80.0,
	}

	if err := policy.VerifyMaxAllowableQuantity(area, 15.0); err != nil {
		t.Errorf("expected 15.0 kg load to be allowed, got error: %v", err)
	}

	if err := policy.VerifyMaxAllowableQuantity(area, 25.0); err == nil {
		t.Error("expected 25.0 kg load to be rejected as it exceeds capacity")
	}
}

func TestExceedsOSHAPSMThreshold(t *testing.T) {
	restrictedChem := &chemical.Chemical{IsRestricted: true}
	generalChem := &chemical.Chemical{IsRestricted: false}

	if !policy.ExceedsOSHAPSMThreshold(restrictedChem, 5000.0) {
		t.Error("expected 5000 kg of restricted chemical to exceed OSHA TQ (4500 kg)")
	}

	if policy.ExceedsOSHAPSMThreshold(generalChem, 8000.0) {
		t.Error("expected 8000 kg of general chemical to not exceed general OSHA TQ (10000 kg)")
	}
}
