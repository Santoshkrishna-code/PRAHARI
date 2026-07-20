package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"

	trainingDomain "prahari/services/training/internal/domain/training"
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
	return fmt.Sprintf("training:detail:%s", id)
}

// Get returns cached record.
func (c *Cache) Get(ctx context.Context, id string) (*trainingDomain.Training, error) {
	data, err := c.client.UniversalClient.Get(ctx, key(id)).Result()
	if err != nil {
		return nil, nil // Cache miss
	}

	var m trainingDomain.Training
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// Set cache details.
func (c *Cache) Set(ctx context.Context, t *trainingDomain.Training) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return c.client.UniversalClient.Set(ctx, key(t.ID), string(data), DefaultTTL).Err()
}

// Invalidate removes cached values.
func (c *Cache) Invalidate(ctx context.Context, id string) error {
	return c.client.UniversalClient.Del(ctx, key(id)).Err()
}
