package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"

	permitDomain "prahari/services/permit/internal/domain/permit"
)

// DefaultTTL sets cache life.
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
	return fmt.Sprintf("permit:detail:%s", id)
}

// Get returns cached value.
func (c *Cache) Get(ctx context.Context, id string) (*permitDomain.Permit, error) {
	data, err := c.client.Get(ctx, key(id))
	if err != nil {
		return nil, nil // Cache miss
	}

	var p permitDomain.Permit
	if err := json.Unmarshal([]byte(data), &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// Set stores value.
func (c *Cache) Set(ctx context.Context, p *permitDomain.Permit) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key(p.ID), string(data), DefaultTTL)
}

// Invalidate clears cache.
func (c *Cache) Invalidate(ctx context.Context, id string) error {
	return c.client.Del(ctx, key(id))
}
