package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	redisSDK "github.com/redis/go-redis/v9"
)

// MSet serializes and writes multiple key-value pairs to Redis using pipelines.
func (c *Cache) MSet(ctx context.Context, pairs map[string]interface{}, ttl time.Duration) error {
	if c.client == nil || c.client.UniversalClient == nil {
		return fmt.Errorf("redis client is uninitialized")
	}

	pipe := c.client.UniversalClient.Pipeline()

	for key, val := range pairs {
		dataBytes, err := c.serializer.Marshal(val)
		if err != nil {
			return fmt.Errorf("failed to marshal batch value for key %s: %w", key, err)
		}
		pipe.Set(ctx, key, dataBytes, ttl)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to execute MSet pipeline: %w", err)
	}

	return nil
}

// MGet fetches multiple keys using pipelines and unmarshals them into the corresponding targets.
// The slice length of keys and targets must be identical.
func (c *Cache) MGet(ctx context.Context, keys []string, targets []interface{}) error {
	if len(keys) != len(targets) {
		return fmt.Errorf("keys and targets slices must have equal lengths")
	}

	if c.client == nil || c.client.UniversalClient == nil {
		return fmt.Errorf("redis client is uninitialized")
	}

	pipe := c.client.UniversalClient.Pipeline()
	cmds := make([]*redisSDK.StringCmd, len(keys))

	for i, key := range keys {
		cmds[i] = pipe.Get(ctx, key)
	}

	_, err := pipe.Exec(ctx)
	if err != nil && !errors.Is(err, redisSDK.Nil) {
		return fmt.Errorf("failed to execute MGet pipeline: %w", err)
	}

	for i, cmd := range cmds {
		dataBytes, err := cmd.Bytes()
		if err != nil {
			if errors.Is(err, redisSDK.Nil) {
				targets[i] = nil // Missed key
				continue
			}
			return fmt.Errorf("failed to read result for key %s: %w", keys[i], err)
		}

		err = c.serializer.Unmarshal(dataBytes, targets[i])
		if err != nil {
			return fmt.Errorf("failed to unmarshal result for key %s: %w", keys[i], err)
		}
	}

	return nil
}
