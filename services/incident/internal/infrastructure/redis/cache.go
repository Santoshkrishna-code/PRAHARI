package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"

	incidentDomain "prahari/services/incident/internal/domain/incident"
)

// DefaultTTL defines the standard cache expiration for incident details.
const DefaultTTL = 15 * time.Minute

// Cache provides Redis-backed caching for incident detail lookups.
type Cache struct {
	client *prahariRedis.Client
}

// NewCache constructs a Cache.
func NewCache(client *prahariRedis.Client) *Cache {
	return &Cache{client: client}
}

// cacheKey generates a namespaced cache key for an incident.
func cacheKey(id string) string {
	return fmt.Sprintf("incident:detail:%s", id)
}

// Get retrieves a cached incident by ID. Returns nil if not cached.
func (c *Cache) Get(ctx context.Context, id string) (*incidentDomain.Incident, error) {
	data, err := c.client.Get(ctx, cacheKey(id))
	if err != nil {
		return nil, nil // Cache miss
	}

	var inc incidentDomain.Incident
	if err := json.Unmarshal([]byte(data), &inc); err != nil {
		return nil, fmt.Errorf("redis: failed to deserialize cached incident: %w", err)
	}
	return &inc, nil
}

// Set stores an incident in the cache with the default TTL.
func (c *Cache) Set(ctx context.Context, inc *incidentDomain.Incident) error {
	data, err := json.Marshal(inc)
	if err != nil {
		return fmt.Errorf("redis: failed to serialize incident for cache: %w", err)
	}

	return c.client.Set(ctx, cacheKey(inc.ID), string(data), DefaultTTL)
}

// Invalidate removes an incident from the cache after a mutation.
func (c *Cache) Invalidate(ctx context.Context, id string) error {
	return c.client.Del(ctx, cacheKey(id))
}

// IncrementCounter increments the hot-incident counter for real-time dashboards.
func (c *Cache) IncrementCounter(ctx context.Context, counterName string) error {
	key := fmt.Sprintf("incident:counter:%s", counterName)
	return c.client.Incr(ctx, key)
}
