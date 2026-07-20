package retry

import (
	"context"
	"time"

	prahariLogger "prahari/shared/logger"
)

// Backoff executes callbacks retrying on transient failures.
type Backoff struct {
	maxAttempts int
	interval    time.Duration
}

// NewBackoff constructs a Backoff helper.
func NewBackoff(attempts int, delay time.Duration) *Backoff {
	return &Backoff{
		maxAttempts: attempts,
		interval:    delay,
	}
}

// Execute retries functions until success or limit thresholds are violated.
func (b *Backoff) Execute(ctx context.Context, fn func() error) error {
	var err error
	currentDelay := b.interval

	for attempt := 1; attempt <= b.maxAttempts; attempt++ {
		err = fn()
		if err == nil {
			return nil
		}

		prahariLogger.Info(ctx, "Transient message delivery failure. Retrying...",
			prahariLogger.Int("attempt", attempt),
			prahariLogger.Err(err))

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(currentDelay):
			// Exponential backoff multiplier
			currentDelay *= 2
		}
	}

	return err
}
