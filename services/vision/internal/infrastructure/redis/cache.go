package redis

import (
	"context"
	"fmt"

	prahariRedis "prahari/shared/redis"
)

type Cache struct {
	client *prahariRedis.Client
}

func NewCache(client *prahariRedis.Client) *Cache {
	return &Cache{client: client}
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	if c.client == nil || c.client.UniversalClient == nil {
		return "", nil
	}
	val, err := c.client.UniversalClient.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("redis get error for key %s: %w", key, err)
	}
	return val, nil
}

func (c *Cache) Set(ctx context.Context, key string, value string) error {
	if c.client == nil || c.client.UniversalClient == nil {
		return nil
	}
	if err := c.client.UniversalClient.Set(ctx, key, value, 0).Err(); err != nil {
		return fmt.Errorf("redis set error for key %s: %w", key, err)
	}
	return nil
}
