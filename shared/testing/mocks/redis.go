package mocks

import (
	"context"
	"time"
)

// MockRedisClient exposes hook parameters to override caching transactions.
type MockRedisClient struct {
	GetFunc func(ctx context.Context, key string) (string, error)
	SetFunc func(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	DelFunc func(ctx context.Context, key string) error
}

// Get delegates transaction to GetFunc.
func (m *MockRedisClient) Get(ctx context.Context, key string) (string, error) {
	if m.GetFunc != nil {
		return m.GetFunc(ctx, key)
	}
	return "", nil
}

// Set delegates transaction to SetFunc.
func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	if m.SetFunc != nil {
		return m.SetFunc(ctx, key, value, ttl)
	}
	return nil
}

// Del delegates transaction to DelFunc.
func (m *MockRedisClient) Del(ctx context.Context, key string) error {
	if m.DelFunc != nil {
		return m.DelFunc(ctx, key)
	}
	return nil
}
