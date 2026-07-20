package cache

import (
	"context"
	"errors"
	"fmt"

	redisSDK "github.com/redis/go-redis/v9"
	prahariRedis "prahari/shared/redis"
)

var (
	// ErrCacheMiss is returned when the requested key is not found in Redis.
	ErrCacheMiss = errors.New("cache: key not found in storage")
)

// Cache coordinates distributed caching operations.
type Cache struct {
	client     *prahariRedis.Client
	serializer prahariRedis.Serializer
}

// NewCache constructs a new Cache instance.
func NewCache(client *prahariRedis.Client, ser prahariRedis.Serializer) *Cache {
	return &Cache{
		client:     client,
		serializer: ser,
	}
}

// Get retrieves a key, deserializes bytes, and unmarshals it into the target destination.
func (c *Cache) Get(ctx context.Context, key string, target interface{}) error {
	if c.client == nil || c.client.UniversalClient == nil {
		return fmt.Errorf("redis client is uninitialized")
	}

	dataBytes, err := c.client.UniversalClient.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redisSDK.Nil) {
			return ErrCacheMiss
		}
		return fmt.Errorf("failed to fetch from Redis cache: %w", err)
	}

	err = c.serializer.Unmarshal(dataBytes, target)
	if err != nil {
		return fmt.Errorf("failed to deserialize cache value: %w", err)
	}

	return nil
}
