package locks

import (
	"context"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"
)

// Redlock implements distributed locks across multiple independent Redis nodes.
type Redlock struct {
	clients []*prahariRedis.Client
	key     string
	value   string
	ttl     time.Duration
	quorum  int
}

// NewRedlock constructs a Redlock coordinator.
func NewRedlock(clients []*prahariRedis.Client, key, value string, ttl time.Duration) (*Redlock, error) {
	if len(clients) == 0 {
		return nil, fmt.Errorf("redlock requires at least one Redis client")
	}

	quorum := len(clients)/2 + 1
	return &Redlock{
		clients: clients,
		key:     key,
		value:   value,
		ttl:     ttl,
		quorum:  quorum,
	}, nil
}

// Lock attempts to acquire the lock on all nodes. Returns error if quorum is not reached.
func (r *Redlock) Lock(ctx context.Context) error {
	start := time.Now()
	acquired := 0

	// Track successfully locked nodes to release on fail
	lockedMutexes := make([]*Mutex, 0, len(r.clients))

	for _, client := range r.clients {
		m := NewMutex(client, r.key, r.value, r.ttl)
		if err := m.Lock(ctx); err == nil {
			acquired++
			lockedMutexes = append(lockedMutexes, m)
		}
	}

	elapsed := time.Since(start)
	// Validity time: TTL minus network drift offset (drift: 2% of TTL + 2ms)
	drift := time.Duration(int64(r.ttl)/50) + 2*time.Millisecond
	validity := r.ttl - elapsed - drift

	// If quorum is reached and validity remains positive, lock is successful
	if acquired >= r.quorum && validity > 0 {
		return nil
	}

	// Quorum not reached or lock expired: release all acquired locks
	for _, m := range lockedMutexes {
		_ = m.Unlock(ctx)
	}

	return ErrLockAcquireFailed
}

// Unlock releases the lock on all nodes.
func (r *Redlock) Unlock(ctx context.Context) error {
	var lastErr error
	for _, client := range r.clients {
		m := NewMutex(client, r.key, r.value, r.ttl)
		// Release locks on all nodes, continuing even if one fails
		if err := m.Unlock(ctx); err != nil {
			lastErr = err
		}
	}
	return lastErr
}
