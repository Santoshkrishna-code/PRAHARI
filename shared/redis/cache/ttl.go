package cache

import (
	"context"
	"fmt"
	"time"
)

// GetTTL returns the remaining TTL duration for a cache key.
func (c *Cache) GetTTL(ctx context.Context, key string) (time.Duration, error) {
	if c.client == nil || c.client.UniversalClient == nil {
		return 0, fmt.Errorf("redis client is uninitialized")
	}

	ttl, err := c.client.UniversalClient.TTL(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to fetch TTL from Redis cache: %w", err)
	}

	// Redis returns -2 if key does not exist
	if ttl == -2*time.Second {
		return 0, ErrCacheMiss
	}

	return ttl, nil
}
