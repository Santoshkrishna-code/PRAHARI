package policy

import (
	"testing"
	"time"

	"prahari/services/document/internal/domain/document"
)

func TestIsReviewOverdue(t *testing.T) {
	past := time.Now().Add(-24 * time.Hour)
	future := time.Now().Add(24 * time.Hour)

	docPast := &document.Document{NextReviewAt: &past}
	docFuture := &document.Document{NextReviewAt: &future}

	if !IsReviewOverdue(docPast) {
		t.Errorf("Expected IsReviewOverdue to be true for past review date")
	}
	if IsReviewOverdue(docFuture) {
		t.Errorf("Expected IsReviewOverdue to be false for future review date")
	}
}

func TestValidateCheckoutLock(t *testing.T) {
	doc := &document.Document{CheckedOutBy: "user-123"}

	if !ValidateCheckoutLock(doc, "user-123") {
		t.Errorf("Expected user-123 to match checkout lock")
	}
	if ValidateCheckoutLock(doc, "user-999") {
		t.Errorf("Expected user-999 to NOT match checkout lock")
	}
}
