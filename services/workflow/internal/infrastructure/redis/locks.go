package redis

import (
	"context"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"
)

// LockManager manages distributed key locks.
type LockManager struct {
	client *prahariRedis.Client
}

// NewLockManager constructs a LockManager.
func NewLockManager(client *prahariRedis.Client) *LockManager {
	return &LockManager{client: client}
}

// AcquireLock sets temporary keys to serialize workflow transitions.
func (m *LockManager) AcquireLock(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	if m.client == nil || m.client.UniversalClient == nil {
		return false, fmt.Errorf("redis client is uninitialized")
	}

	ok, err := m.client.UniversalClient.SetNX(ctx, "lock:"+key, "active", ttl).Result()
	return ok, err
}
