package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Client wraps the go-redis client to perform pool operations.
type Client struct {
	UniversalClient redis.UniversalClient
}

// NewClient bootstraps a connection pool to Redis.
func NewClient(cfg Config) (*Client, error) {
	// Standard configuration mapper
	opts := &redis.Options{
		Addr:       cfg.Address,
		Password:   cfg.Password,
		DB:         cfg.DB,
		PoolSize:   cfg.PoolSize,
		MaxRetries: cfg.MaxRetries,
	}

	rClient := redis.NewClient(opts)

	// Validate connection on startup (fail-fast rule)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := rClient.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis at %s: %w", cfg.Address, err)
	}

	return &Client{
		UniversalClient: rClient,
	}, nil
}

// Close closes the underlying Redis client pools.
func (c *Client) Close() error {
	if c.UniversalClient == nil {
		return nil
	}
	return c.UniversalClient.Close()
}
