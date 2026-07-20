package redis

import (
	"context"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"
)

// Cache adapter executing Redis client commands.
type Cache struct {
	client *prahariRedis.Client
}

// NewCache constructs a Cache.
func NewCache(client *prahariRedis.Client) *Cache {
	return &Cache{client: client}
}

// BlacklistToken flags access tokens invalid until expiration limits.
func (c *Cache) BlacklistToken(ctx context.Context, jti string, ttl time.Duration) error {
	if c.client == nil || c.client.UniversalClient == nil {
		return fmt.Errorf("redis client is uninitialized")
	}

	key := fmt.Sprintf("blacklist:token:%s", jti)
	return c.client.UniversalClient.Set(ctx, key, "revoked", ttl).Err()
}
