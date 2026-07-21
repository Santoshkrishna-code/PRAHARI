package bulkhead

import (
	"context"
	"errors"
)

var (
	// ErrBulkheadFull is returned when concurrent execution limit is reached.
	ErrBulkheadFull = errors.New("bulkhead: concurrent execution queue is full")
)

// Bulkhead controls resource allocation by restricting concurrent executions.
type Bulkhead struct {
	sem chan struct{}
}

// NewBulkhead constructs a Bulkhead with the specified capacity limit.
func NewBulkhead(maxConcurrent int) *Bulkhead {
	return &Bulkhead{
		sem: make(chan struct{}, maxConcurrent),
	}
}

// Execute runs the function only if a semaphore slot is available, releasing it on completion.
func (b *Bulkhead) Execute(ctx context.Context, fn func() error) error {
	if b.sem == nil {
		return errors.New("bulkhead semaphore is uninitialized")
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case b.sem <- struct{}{}:
		// Claimed slot: execute function and release on defer exit
		defer func() {
			<-b.sem
		}()
		return fn()
	default:
		// Queue full
		return ErrBulkheadFull
	}
}
