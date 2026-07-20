package cache

import (
	"context"
	"fmt"
	"time"
)

// Set serializes the value and writes it to Redis with the specified TTL.
func (c *Cache) Set(ctx context.Context, key string, val interface{}, ttl time.Duration) error {
	if c.client == nil || c.client.UniversalClient == nil {
		return fmt.Errorf("redis client is uninitialized")
	}

	dataBytes, err := c.serializer.Marshal(val)
	if err != nil {
		return fmt.Errorf("failed to serialize cache value: %w", err)
	}

	err = c.client.UniversalClient.Set(ctx, key, dataBytes, ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to write to Redis cache: %w", err)
	}

	return nil
}
