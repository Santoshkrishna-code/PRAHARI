package middleware

import (
	"context"
	"fmt"
	"sync/atomic"
)

// MetricsTracker aggregates atomic counters for topic transactions.
type MetricsTracker struct {
	SuccessCount int64
	FailureCount int64
}

// MetricsMiddleware increments success and failure counts depending on handler exits.
func (t *MetricsTracker) MetricsMiddleware(handler func(ctx context.Context, key []byte, val []byte) error) func(ctx context.Context, key []byte, val []byte) error {
	return func(ctx context.Context, key, val []byte) error {
		err := handler(ctx, key, val)
		if err != nil {
			atomic.AddInt64(&t.FailureCount, 1)
			fmt.Printf("[KAFKA-MID] Metrics: incremented failure count to %d\n", atomic.LoadInt64(&t.FailureCount))
		} else {
			atomic.AddInt64(&t.SuccessCount, 1)
			fmt.Printf("[KAFKA-MID] Metrics: incremented success count to %d\n", atomic.LoadInt64(&t.SuccessCount))
		}
		return err
	}
}
