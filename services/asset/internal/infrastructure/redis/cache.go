package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"

	assetDomain "prahari/services/asset/internal/domain/asset"
)

// DefaultTTL cache length (15m).
const DefaultTTL = 15 * time.Minute

// Cache implements cache operations backed by Redis.
type Cache struct {
	client *prahariRedis.Client
}

// NewCache instantiates Cache.
func NewCache(client *prahariRedis.Client) *Cache {
	return &Cache{client: client}
}

// key helper.
func key(id string) string {
	return fmt.Sprintf("asset:detail:%s", id)
}

// Get returns cached record.
func (c *Cache) Get(ctx context.Context, id string) (*assetDomain.Asset, error) {
	data, err := c.client.Get(ctx, key(id))
	if err != nil {
		return nil, nil // Cache miss
	}

	var a assetDomain.Asset
	if err := json.Unmarshal([]byte(data), &a); err != nil {
		return nil, err
	}
	return &a, nil
}

// Set cache details.
func (c *Cache) Set(ctx context.Context, a *assetDomain.Asset) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key(a.ID), string(data), DefaultTTL)
}

// Invalidate removes cached values.
func (c *Cache) Invalidate(ctx context.Context, id string) error {
	return c.client.Del(ctx, key(id))
}
