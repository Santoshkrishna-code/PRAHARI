package retry

import (
	"context"
	"fmt"
	"time"
)

// Retry executes a callback handler, retrying up to maxAttempts on failure.
func Retry(ctx context.Context, b Backoff, maxAttempts int, fn func() error) error {
	var lastErr error

	for attempt := 0; attempt < maxAttempts; attempt++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			err := fn()
			if err == nil {
				return nil
			}

			lastErr = err

			// Backoff before next attempt
			sleepDur := b.Duration(attempt)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(sleepDur):
			}
		}
	}

	return fmt.Errorf("resilience: retry process exhausted after %d attempts: %w", maxAttempts, lastErr)
}
