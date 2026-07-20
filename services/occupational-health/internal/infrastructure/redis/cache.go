package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"

	profileDomain "prahari/services/occupational-health/internal/domain/healthprofile"
)

const DefaultTTL = 15 * time.Minute

type Cache struct {
	client *prahariRedis.Client
}

func NewCache(client *prahariRedis.Client) *Cache {
	return &Cache{client: client}
}

func key(id string) string {
	return fmt.Sprintf("health:profile:detail:%s", id)
}

func (c *Cache) Get(ctx context.Context, id string) (*profileDomain.HealthProfile, error) {
	data, err := c.client.UniversalClient.Get(ctx, key(id)).Result()
	if err != nil {
		return nil, nil // Cache miss
	}

	var m profileDomain.HealthProfile
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (c *Cache) Set(ctx context.Context, p *profileDomain.HealthProfile) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return c.client.UniversalClient.Set(ctx, key(p.ID), string(data), DefaultTTL).Err()
}

func (c *Cache) Invalidate(ctx context.Context, id string) error {
	return c.client.UniversalClient.Del(ctx, key(id)).Err()
}
