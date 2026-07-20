package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"

	auditDomain "prahari/services/audit/internal/domain/audit"
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
	return fmt.Sprintf("audit:detail:%s", id)
}

// Get returns cached record.
func (c *Cache) Get(ctx context.Context, id string) (*auditDomain.Audit, error) {
	data, err := c.client.Get(ctx, key(id))
	if err != nil {
		return nil, nil // Cache miss
	}

	var m auditDomain.Audit
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// Set cache details.
func (c *Cache) Set(ctx context.Context, m *auditDomain.Audit) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key(m.ID), string(data), DefaultTTL)
}

// Invalidate removes cached values.
func (c *Cache) Invalidate(ctx context.Context, id string) error {
	return c.client.Del(ctx, key(id))
}
}
