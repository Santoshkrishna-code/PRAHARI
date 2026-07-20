package cache

import (
	"context"
	"fmt"
)

// Delete evicts a key from the Redis cache.
func (c *Cache) Delete(ctx context.Context, key string) error {
	if c.client == nil || c.client.UniversalClient == nil {
		return fmt.Errorf("redis client is uninitialized")
	}

	err := c.client.UniversalClient.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to evict from Redis cache: %w", err)
	}

	return nil
}
