package repository

import (
	"context"
	"sync"
	"time"

	"prahari/services/auth-service/internal/domain"
)

type cacheItem struct {
	user      *domain.User
	expiresAt time.Time
}

type InMemoryUserCache struct {
	mu    sync.RWMutex
	items map[string]cacheItem
	ttl   time.Duration
}

func NewInMemoryUserCache(ttl time.Duration) UserCache {
	cache := &InMemoryUserCache{
		items: make(map[string]cacheItem),
		ttl:   ttl,
	}
	
	// Start clean-up routine
	go cache.startCleanupRoutine(1 * time.Minute)
	return cache
}

func (c *InMemoryUserCache) Get(ctx context.Context, tokenStr string) (*domain.User, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	item, found := c.items[tokenStr]
	if !found || time.Now().After(item.expiresAt) {
		return nil, domain.ErrUserNotFound
	}
	
	return item.user, nil
}

func (c *InMemoryUserCache) Set(ctx context.Context, tokenStr string, user *domain.User) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.items[tokenStr] = cacheItem{
		user:      user,
		expiresAt: time.Now().Add(c.ttl),
	}
	return nil
}

func (c *InMemoryUserCache) Delete(ctx context.Context, tokenStr string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	delete(c.items, tokenStr)
	return nil
}

func (c *InMemoryUserCache) startCleanupRoutine(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, item := range c.items {
			if now.After(item.expiresAt) {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}
