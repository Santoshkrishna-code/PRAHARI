package policy_test

import (
	"testing"
	"time"

	"prahari/services/integration/internal/domain/policy"
	"prahari/services/integration/internal/domain/retry"
)

func TestCalculateNextBackoff(t *testing.T) {
	d1 := policy.CalculateNextBackoff(1)
	if d1 != 10*time.Second {
		t.Errorf("expected 10 seconds backoff for retry 1, got %v", d1)
	}

	d2 := policy.CalculateNextBackoff(2)
	if d2 != 20*time.Second {
		t.Errorf("expected 20 seconds backoff for retry 2, got %v", d2)
	}
}

func TestIsRetryExhausted(t *testing.T) {
	mActive := &retry.Message{RetryCount: 2, MaxRetries: 3}
	mExhausted := &retry.Message{RetryCount: 3, MaxRetries: 3}

	if policy.IsRetryExhausted(mActive) {
		t.Error("expected active retry message not to be exhausted")
	}

	if !policy.IsRetryExhausted(mExhausted) {
		t.Error("expected exhausted retry message to return true")
	}
}
