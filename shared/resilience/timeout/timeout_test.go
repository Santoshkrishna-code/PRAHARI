package timeout_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"prahari/shared/resilience/timeout"
)

func TestTimeout_Trigger(t *testing.T) {
	ctx := context.Background()

	// Case 1: Succeeds before timeout
	err := timeout.Execute(ctx, 50*time.Millisecond, func(ctx context.Context) error {
		time.Sleep(10 * time.Millisecond)
		return nil
	})
	if err != nil {
		t.Fatalf("expected successful execution, got: %v", err)
	}

	// Case 2: Blocked and times out -> expect context.DeadlineExceeded
	err = timeout.Execute(ctx, 10*time.Millisecond, func(ctx context.Context) error {
		time.Sleep(50 * time.Millisecond)
		return nil
	})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("expected context.DeadlineExceeded, got %v", err)
	}
}
