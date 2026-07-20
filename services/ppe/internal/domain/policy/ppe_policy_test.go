package policy

import (
	"testing"
	"time"

	"prahari/services/ppe/internal/domain/inventory"
	"prahari/services/ppe/internal/domain/ppeitem"
)

func TestIsPPEExpired(t *testing.T) {
	past := time.Now().Add(-24 * time.Hour)
	future := time.Now().Add(24 * time.Hour)

	itemPast := &ppeitem.Item{ExpiryDate: past}
	itemFuture := &ppeitem.Item{ExpiryDate: future}

	if !IsPPEExpired(itemPast) {
		t.Errorf("Expected expired item to return true")
	}
	if IsPPEExpired(itemFuture) {
		t.Errorf("Expected non-expired item to return false")
	}
}

func TestIsLowStock(t *testing.T) {
	stockLow := &inventory.Stock{QuantityOnHand: 2, BufferLevel: 5}
	stockGood := &inventory.Stock{QuantityOnHand: 10, BufferLevel: 5}

	if !IsLowStock(stockLow) {
		t.Errorf("Expected IsLowStock to return true for stock level below buffer")
	}
	if IsLowStock(stockGood) {
		t.Errorf("Expected IsLowStock to return false for stock level above buffer")
	}
}

func TestCanAssignToUser(t *testing.T) {
	future := time.Now().Add(24 * time.Hour)
	itemOk := &ppeitem.Item{Status: "AVAILABLE", ExpiryDate: future}
	itemNotAvailable := &ppeitem.Item{Status: "ISSUED", ExpiryDate: future}

	if !CanAssignToUser(itemOk) {
		t.Errorf("Expected CanAssignToUser to return true for available non-expired item")
	}
	if CanAssignToUser(itemNotAvailable) {
		t.Errorf("Expected CanAssignToUser to return false for already issued item")
	}
}
