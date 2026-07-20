package policy

import (
	"time"

	"prahari/services/document/internal/domain/document"
)

// IsReviewOverdue checks if periodic SME document review is overdue.
func IsReviewOverdue(doc *document.Document) bool {
	if doc.NextReviewAt == nil {
		return false
	}
	return time.Now().After(*doc.NextReviewAt)
}

// CalculateNextReviewDate computes next periodic review deadline based on review cycle months.
func CalculateNextReviewDate(from time.Time, cycleMonths int) time.Time {
	if cycleMonths <= 0 {
		cycleMonths = 12 // Default 1 year review cycle
	}
	return from.AddDate(0, cycleMonths, 0)
}

// ValidateCheckoutLock ensures document is not checked out by another user.
func ValidateCheckoutLock(doc *document.Document, userID string) bool {
	if doc.CheckedOutBy == "" {
		return true
	}
	return doc.CheckedOutBy == userID
}
