package ratelimit

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	prahariRedis "prahari/shared/redis"
)

// SlidingWindowLimiter implements a sliding-window log rate limiter using ZSETs.
type SlidingWindowLimiter struct {
	client *prahariRedis.Client
	limit  int64
	window time.Duration
}

// NewSlidingWindowLimiter constructs a new SlidingWindowLimiter.
func NewSlidingWindowLimiter(client *prahariRedis.Client, limit int64, window time.Duration) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		client: client,
		limit:  limit,
		window: window,
	}
}

// Allow returns true if the request for the key is within limit bounds over the sliding window.
func (l *SlidingWindowLimiter) Allow(ctx context.Context, key string) (bool, error) {
	if l.client == nil || l.client.UniversalClient == nil {
		return false, fmt.Errorf("redis client is uninitialized")
	}

	now := time.Now()
	nowMs := now.UnixNano() / int64(time.Millisecond)
	clearBeforeMs := nowMs - (l.window.Nanoseconds() / int64(time.Millisecond))
	redisKey := fmt.Sprintf("ratelimit:sliding:%s", key)

	// Unique member value to prevent overlapping set values
	uniqueMember := strconv.FormatInt(now.UnixNano(), 10)

	pipe := l.client.UniversalClient.Pipeline()
	
	// Remove events outside current sliding window
	pipe.ZRemRangeByScore(ctx, redisKey, "0", strconv.FormatInt(clearBeforeMs, 10))
	
	// Query current window count
	cardCmd := pipe.ZCard(ctx, redisKey)
	
	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to run sliding window check pipeline: %w", err)
	}

	currentRequests, err := cardCmd.Result()
	if err != nil {
		return false, fmt.Errorf("failed to fetch sliding window card result: %w", err)
	}

	if currentRequests >= l.limit {
		return false, nil
	}

	// Insert current request and extend key TTL
	pipe = l.client.UniversalClient.Pipeline()
	pipe.ZAdd(ctx, redisKey, redis.Z{
		Score:  float64(nowMs),
		Member: uniqueMember,
	})
	pipe.Expire(ctx, redisKey, l.window)

	_, err = pipe.Exec(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to insert sliding window record: %w", err)
	}

	return true, nil
}
