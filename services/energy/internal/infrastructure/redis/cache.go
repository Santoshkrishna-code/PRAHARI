package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"

	profileDomain "prahari/services/energy/internal/domain/energyprofile"
)

const DefaultTTL = 15 * time.Minute

type Cache struct {
	client *prahariRedis.Client
}

func NewCache(client *prahariRedis.Client) *Cache {
	return &Cache{client: client}
}

func key(id string) string {
	return fmt.Sprintf("energy:profile:detail:%s", id)
}

func (c *Cache) Get(ctx context.Context, id string) (*profileDomain.Profile, error) {
	data, err := c.client.UniversalClient.Get(ctx, key(id)).Result()
	if err != nil {
		return nil, nil // Cache miss
	}

	var p profileDomain.Profile
	if err := json.Unmarshal([]byte(data), &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (c *Cache) Set(ctx context.Context, p *profileDomain.Profile) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return c.client.UniversalClient.Set(ctx, key(p.ID), string(data), DefaultTTL).Err()
}

func (c *Cache) Invalidate(ctx context.Context, id string) error {
	return c.client.UniversalClient.Del(ctx, key(id)).Err()
}
