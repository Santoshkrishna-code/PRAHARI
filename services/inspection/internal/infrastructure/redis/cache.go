package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"

	inspectionDomain "prahari/services/inspection/internal/domain/inspection"
)

// DefaultTTL sets cache lifespan rules.
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
	return fmt.Sprintf("inspection:detail:%s", id)
}

// Get returns cached record.
func (c *Cache) Get(ctx context.Context, id string) (*inspectionDomain.Inspection, error) {
	data, err := c.client.Get(ctx, key(id))
	if err != nil {
		return nil, nil // Cache miss
	}

	var i inspectionDomain.Inspection
	if err := json.Unmarshal([]byte(data), &i); err != nil {
		return nil, err
	}
	return &i, nil
}

// Set cache details.
func (c *Cache) Set(ctx context.Context, i *inspectionDomain.Inspection) error {
	data, err := json.Marshal(i)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key(i.ID), string(data), DefaultTTL)
}

// Invalidate removes cached values.
func (c *Cache) Invalidate(ctx context.Context, id string) error {
	return c.client.Del(ctx, key(id))
}
