package policy

import (
	"time"

	"prahari/services/ppe/internal/domain/inventory"
	"prahari/services/ppe/internal/domain/ppeitem"
)

// IsPPEExpired checks if a PPE item's expiry date has passed.
func IsPPEExpired(item *ppeitem.Item) bool {
	return time.Now().After(item.ExpiryDate)
}

// IsLowStock checks if the stock level of a PPE model has dropped below safety buffer levels.
func IsLowStock(stock *inventory.Stock) bool {
	return stock.QuantityOnHand <= stock.BufferLevel
}

// CanAssignToUser verifies standard parameters like item eligibility before issuing.
func CanAssignToUser(item *ppeitem.Item) bool {
	return item.Status == "AVAILABLE" && !IsPPEExpired(item)
}
