package timeout

import (
	"context"
	"time"
)

// Execute runs the function wrapped with a timeout context.
// Returns context.DeadlineExceeded if the execution time limit is crossed.
func Execute(ctx context.Context, limit time.Duration, fn func(ctx context.Context) error) error {
	tCtx, cancel := context.WithTimeout(ctx, limit)
	defer cancel()

	ch := make(chan error, 1)
	go func() {
		ch <- fn(tCtx)
	}()

	select {
	case <-tCtx.Done():
		return tCtx.Err()
	case err := <-ch:
		return err
	}
}
