package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"

	aspectDomain "prahari/services/environmental/internal/domain/environment"
)

const DefaultTTL = 15 * time.Minute

type Cache struct {
	client *prahariRedis.Client
}

func NewCache(client *prahariRedis.Client) *Cache {
	return &Cache{client: client}
}

func key(id string) string {
	return fmt.Sprintf("environmental:aspect:detail:%s", id)
}

func (c *Cache) Get(ctx context.Context, id string) (*aspectDomain.EnvironmentalAspect, error) {
	data, err := c.client.UniversalClient.Get(ctx, key(id)).Result()
	if err != nil {
		return nil, nil // Cache miss
	}

	var m aspectDomain.EnvironmentalAspect
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (c *Cache) Set(ctx context.Context, a *aspectDomain.EnvironmentalAspect) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}
	return c.client.UniversalClient.Set(ctx, key(a.ID), string(data), DefaultTTL).Err()
}

func (c *Cache) Invalidate(ctx context.Context, id string) error {
	return c.client.UniversalClient.Del(ctx, key(id)).Err()
}
