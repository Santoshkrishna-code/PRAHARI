package ratelimit

import (
	"context"
	"fmt"
	"time"

	prahariRedis "prahari/shared/redis"
)

// TokenBucketLimiter implements a token bucket rate limiter using Redis hashes and Lua scripts.
type TokenBucketLimiter struct {
	client   *prahariRedis.Client
	capacity int64
	fillRate float64 // tokens per second
}

// NewTokenBucketLimiter constructs a new TokenBucketLimiter.
func NewTokenBucketLimiter(client *prahariRedis.Client, capacity int64, fillRate float64) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		client:   client,
		capacity: capacity,
		fillRate: fillRate,
	}
}

// Allow returns true if a token was successfully consumed from the bucket.
func (l *TokenBucketLimiter) Allow(ctx context.Context, key string) (bool, error) {
	if l.client == nil || l.client.UniversalClient == nil {
		return false, fmt.Errorf("redis client is uninitialized")
	}

	const luaScript = `
		local key = KEYS[1]
		local capacity = tonumber(ARGV[1])
		local fill_rate = tonumber(ARGV[2])
		local now = tonumber(ARGV[3])
		local requested = 1

		local state = redis.call("HMGET", key, "tokens", "last_updated")
		local tokens = tonumber(state[1])
		local last_updated = tonumber(state[2])

		if not tokens then
			tokens = capacity
			last_updated = now
		else
			local elapsed = math.max(0, now - last_updated)
			tokens = math.min(capacity, tokens + (elapsed * fill_rate))
			last_updated = now
		end

		if tokens >= requested then
			tokens = tokens - requested
			redis.call("HMSET", key, "tokens", tokens, "last_updated", last_updated)
			redis.call("EXPIRE", key, 86400) -- expire key in 24 hours to prevent memory leak
			return 1
		else
			redis.call("HMSET", key, "tokens", tokens, "last_updated", last_updated)
			return 0
		end
	`

	redisKey := fmt.Sprintf("ratelimit:bucket:%s", key)
	nowSecs := float64(time.Now().UnixNano()) / 1e9

	res, err := l.client.UniversalClient.Eval(
		ctx,
		luaScript,
		[]string{redisKey},
		l.capacity,
		l.fillRate,
		nowSecs,
	).Result()

	if err != nil {
		return false, fmt.Errorf("failed to execute token bucket Lua script: %w", err)
	}

	if val, ok := res.(int64); ok && val == 1 {
		return true, nil
	}

	return false, nil
}
