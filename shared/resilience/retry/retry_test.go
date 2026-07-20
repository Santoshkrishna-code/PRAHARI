package retry_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"prahari/shared/resilience/retry"
)

func TestRetry_SuccessAndExhausted(t *testing.T) {
	b := retry.Backoff{
		Min:    1 * time.Millisecond,
		Max:    5 * time.Millisecond,
		Factor: 2.0,
		Jitter: false,
	}

	ctx := context.Background()

	// Case 1: Succeeds on second try
	attempts := 0
	err := retry.Retry(ctx, b, 3, func() error {
		attempts++
		if attempts < 2 {
			return errors.New("fail")
		}
		return nil
	})

	if err != nil {
		t.Fatalf("expected retry success on second try, got: %v", err)
	}

	if attempts != 2 {
		t.Errorf("expected 2 execution attempts, got %d", attempts)
	}

	// Case 2: Exhausted
	err = retry.Retry(ctx, b, 2, func() error {
		return errors.New("persistent fail")
	})

	if err == nil {
		t.Error("expected retry exhaustion error, got nil")
	}
}
