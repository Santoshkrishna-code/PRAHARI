package pipeline

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	prahariRedis "prahari/shared/redis"
)

// Pipeline wraps the go-redis Pipeline client.
type Pipeline struct {
	pipe redis.Pipeliner
}

// NewPipeline constructs a new Pipeline wrapper.
func NewPipeline(client *prahariRedis.Client) (*Pipeline, error) {
	if client == nil || client.UniversalClient == nil {
		return nil, fmt.Errorf("redis client is uninitialized")
	}
	return &Pipeline{
		pipe: client.UniversalClient.Pipeline(),
	}, nil
}

// Get appends a GET command to the pipeline.
func (p *Pipeline) Get(ctx context.Context, key string) *redis.StringCmd {
	return p.pipe.Get(ctx, key)
}

// Set appends a SET command to the pipeline.
func (p *Pipeline) Set(ctx context.Context, key string, val interface{}, ttl time.Duration) *redis.StatusCmd {
	return p.pipe.Set(ctx, key, val, ttl)
}

// Exec executes all queued pipeline commands.
func (p *Pipeline) Exec(ctx context.Context) ([]redis.Cmder, error) {
	return p.pipe.Exec(ctx)
}
