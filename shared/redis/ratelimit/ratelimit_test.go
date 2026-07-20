package ratelimit_test

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	prahariRedis "prahari/shared/redis"
	"prahari/shared/redis/ratelimit"
)

func TestFixedWindowLimiter(t *testing.T) {
	mr, _ := miniredis.Run()
	defer mr.Close()

	client, _ := prahariRedis.NewClient(prahariRedis.Config{Address: mr.Addr()})
	defer client.Close()

	ctx := context.Background()
	// Limit: 2 requests per 10 seconds
	limiter := ratelimit.NewFixedWindowLimiter(client, 2, 10*time.Second)

	ok, _ := limiter.Allow(ctx, "user-1")
	if !ok {
		t.Error("expected first request to be allowed")
	}

	ok, _ = limiter.Allow(ctx, "user-1")
	if !ok {
		t.Error("expected second request to be allowed")
	}

	ok, _ = limiter.Allow(ctx, "user-1")
	if ok {
		t.Error("expected third request to be blocked")
	}
}

func TestSlidingWindowLimiter(t *testing.T) {
	mr, _ := miniredis.Run()
	defer mr.Close()

	client, _ := prahariRedis.NewClient(prahariRedis.Config{Address: mr.Addr()})
	defer client.Close()

	ctx := context.Background()
	// Limit: 2 requests per 1 second
	limiter := ratelimit.NewSlidingWindowLimiter(client, 2, 1*time.Second)

	ok, _ := limiter.Allow(ctx, "user-2")
	if !ok {
		t.Error("expected request 1 to pass")
	}

	ok, _ = limiter.Allow(ctx, "user-2")
	if !ok {
		t.Error("expected request 2 to pass")
	}

	ok, _ = limiter.Allow(ctx, "user-2")
	if ok {
		t.Error("expected request 3 to be blocked")
	}

	// Sleep to let sliding window age
	time.Sleep(1100 * time.Millisecond)

	ok, _ = limiter.Allow(ctx, "user-2")
	if !ok {
		t.Error("expected request to pass after sliding window expired")
	}
}

func TestTokenBucketLimiter(t *testing.T) {
	mr, _ := miniredis.Run()
	defer mr.Close()

	client, _ := prahariRedis.NewClient(prahariRedis.Config{Address: mr.Addr()})
	defer client.Close()

	ctx := context.Background()
	// Capacity 2, fill rate 1 token/sec
	limiter := ratelimit.NewTokenBucketLimiter(client, 2, 1.0)

	ok, _ := limiter.Allow(ctx, "user-3")
	if !ok {
		t.Error("expected request 1 to consume token")
	}

	ok, _ = limiter.Allow(ctx, "user-3")
	if !ok {
		t.Error("expected request 2 to consume token")
	}

	ok, _ = limiter.Allow(ctx, "user-3")
	if ok {
		t.Error("expected request 3 to fail due to empty bucket")
	}
}
