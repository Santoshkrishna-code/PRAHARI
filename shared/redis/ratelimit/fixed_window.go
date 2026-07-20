package ratelimit

import (
	"context"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"
)

// FixedWindowLimiter implements a fixed-window counter rate limiter.
type FixedWindowLimiter struct {
	client *prahariRedis.Client
	limit  int64
	window time.Duration
}

// NewFixedWindowLimiter constructs a new FixedWindowLimiter.
func NewFixedWindowLimiter(client *prahariRedis.Client, limit int64, window time.Duration) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		client: client,
		limit:  limit,
		window: window,
	}
}

// Allow returns true if the request for the key is within limit bounds during the current window.
func (l *FixedWindowLimiter) Allow(ctx context.Context, key string) (bool, error) {
	if l.client == nil || l.client.UniversalClient == nil {
		return false, fmt.Errorf("redis client is uninitialized")
	}

	now := time.Now()
	// Calculate window bucket index (fixed boundary)
	bucket := now.UnixNano() / l.window.Nanoseconds()
	redisKey := fmt.Sprintf("ratelimit:fixed:%s:%d", key, bucket)

	pipe := l.client.UniversalClient.Pipeline()
	incr := pipe.Incr(ctx, redisKey)
	pipe.Expire(ctx, redisKey, l.window)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to execute fixed window pipeline: %w", err)
	}

	count, err := incr.Result()
	if err != nil {
		return false, fmt.Errorf("failed to fetch incr value: %w", err)
	}

	return count <= l.limit, nil
}
