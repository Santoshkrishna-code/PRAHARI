package redis

import (
	"context"
	"fmt"
)

// Ping validates Redis server connectivity.
func (c *Client) Ping(ctx context.Context) error {
	if c.UniversalClient == nil {
		return fmt.Errorf("redis client is uninitialized")
	}
	return c.UniversalClient.Ping(ctx).Err()
}
