package bootstrap

import (
	"context"
	"fmt"

	prahariRedis "prahari/shared/redis"
)

// InitRedis connects Redis client pathway.
func InitRedis(ctx context.Context, addr string) (*prahariRedis.Client, error) {
	client, err := prahariRedis.NewClient(prahariRedis.Config{
		Address: addr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	return client, nil
}
