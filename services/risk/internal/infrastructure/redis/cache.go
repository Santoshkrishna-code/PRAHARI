package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"

	riskDomain "prahari/services/risk/internal/domain/risk"
)

// DefaultTTL sets cache lifespan.
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
	return fmt.Sprintf("risk:detail:%s", id)
}

// Get returns cached record.
func (c *Cache) Get(ctx context.Context, id string) (*riskDomain.Risk, error) {
	data, err := c.client.Get(ctx, key(id))
	if err != nil {
		return nil, nil // Cache miss
	}

	var r riskDomain.Risk
	if err := json.Unmarshal([]byte(data), &r); err != nil {
		return nil, err
	}
	return &r, nil
}

// Set cache details.
func (c *Cache) Set(ctx context.Context, r *riskDomain.Risk) error {
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key(r.ID), string(data), DefaultTTL)
}

// Invalidate removes cached values.
func (c *Cache) Invalidate(ctx context.Context, id string) error {
	return c.client.Del(ctx, key(id))
}
}
