package policy

import (
	"testing"

	"prahari/services/loto/internal/domain/isolationcertificate"
	"prahari/services/loto/internal/domain/lock"
	"prahari/services/loto/internal/domain/verification"
)

func TestCanCommenceMaintenance(t *testing.T) {
	certOk := &isolationcertificate.Certificate{Status: "ZERO_ENERGY_VERIFIED"}
	certNotOk := &isolationcertificate.Certificate{Status: "PLANNED"}

	vPass := &verification.ZeroEnergy{TestPassed: true}
	vFail := &verification.ZeroEnergy{TestPassed: false}

	if !CanCommenceMaintenance(certOk, vPass) {
		t.Errorf("Expected CanCommenceMaintenance to return true for verified certificate and passed test")
	}
	if CanCommenceMaintenance(certNotOk, vPass) {
		t.Errorf("Expected CanCommenceMaintenance to return false for unverified certificate status")
	}
	if CanCommenceMaintenance(certOk, vFail) {
		t.Errorf("Expected CanCommenceMaintenance to return false for failed verification test")
	}
}

func TestAllLocksRemoved(t *testing.T) {
	lock1 := &lock.Lock{Status: "AVAILABLE"}
	lock2 := &lock.Lock{Status: "AVAILABLE"}
	lock3 := &lock.Lock{Status: "APPLIED"}

	if !AllLocksRemoved([]*lock.Lock{lock1, lock2}) {
		t.Errorf("Expected AllLocksRemoved to return true when all locks are available")
	}
	if AllLocksRemoved([]*lock.Lock{lock1, lock3}) {
		t.Errorf("Expected AllLocksRemoved to return false when a lock is still applied")
	}
}
