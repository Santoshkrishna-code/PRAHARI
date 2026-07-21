package locks

import (
	"context"
	"errors"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"
)

var (
	// ErrLockAcquireFailed is returned if the lock is already held.
	ErrLockAcquireFailed = errors.New("locks: failed to acquire distributed lock")
)

// Mutex represents a single distributed lock instance backed by Redis.
type Mutex struct {
	client *prahariRedis.Client
	key    string
	value  string // unique identifier for the lock holder (e.g. UUID)
	ttl    time.Duration
}

// NewMutex constructs a new Mutex instance.
func NewMutex(client *prahariRedis.Client, key, value string, ttl time.Duration) *Mutex {
	return &Mutex{
		client: client,
		key:    key,
		value:  value,
		ttl:    ttl,
	}
}

// Lock attempts to acquire the distributed lock. Returns ErrLockAcquireFailed if already locked.
func (m *Mutex) Lock(ctx context.Context) error {
	if m.client == nil || m.client.UniversalClient == nil {
		return fmt.Errorf("redis client is uninitialized")
	}

	// Atomically set key if not exists (NX) with expiry (PX)
	ok, err := m.client.UniversalClient.SetNX(ctx, m.key, m.value, m.ttl).Result()
	if err != nil {
		return fmt.Errorf("failed to execute SetNX command: %w", err)
	}

	if !ok {
		return ErrLockAcquireFailed
	}

	return nil
}

// Unlock releases the lock safely using a Lua script to ensure key value validation.
func (m *Mutex) Unlock(ctx context.Context) error {
	if m.client == nil || m.client.UniversalClient == nil {
		return fmt.Errorf("redis client is uninitialized")
	}

	// Lua script: release lock only if key holds our unique value
	const luaScript = `
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else
			return 0
		end
	`

	res, err := m.client.UniversalClient.Eval(ctx, luaScript, []string{m.key}, m.value).Result()
	if err != nil {
		return fmt.Errorf("failed to execute unlock Lua script: %w", err)
	}

	// If script returns 0, the lock has expired or belongs to someone else
	if count, ok := res.(int64); ok && count == 0 {
		return fmt.Errorf("locks: unlock failed: lock expired or owned by another process")
	}

	return nil
}
