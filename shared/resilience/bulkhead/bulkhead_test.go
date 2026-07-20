package bulkhead_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"prahari/shared/resilience/bulkhead"
)

func TestBulkhead_ConcurrencyLimit(t *testing.T) {
	// Limit: 2 concurrent tasks
	b := bulkhead.NewBulkhead(2)
	ctx := context.Background()

	var wg sync.WaitGroup
	wg.Add(2)

	// Block two tasks inside the bulkhead
	go func() {
		_ = b.Execute(ctx, func() error {
			wg.Done()
			time.Sleep(50 * time.Millisecond)
			return nil
		})
	}()

	go func() {
		_ = b.Execute(ctx, func() error {
			wg.Done()
			time.Sleep(50 * time.Millisecond)
			return nil
		})
	}()

	// Wait for both goroutines to enter the bulkhead execution block
	wg.Wait()

	// Third request should be blocked immediately (queue is full) -> expect ErrBulkheadFull
	err := b.Execute(ctx, func() error {
		return nil
	})

	if !errors.Is(err, bulkhead.ErrBulkheadFull) {
		t.Errorf("expected ErrBulkheadFull on third execution, got %v", err)
	}
}
