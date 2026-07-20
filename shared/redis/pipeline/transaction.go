package pipeline

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	prahariRedis "prahari/shared/redis"
)

// Transaction wraps Redis MULTI/EXEC operations.
type Transaction struct {
	client *prahariRedis.Client
}

// NewTransaction constructs a new Transaction wrapper.
func NewTransaction(client *prahariRedis.Client) *Transaction {
	return &Transaction{client: client}
}

// Watch monitors keys for changes and runs the transaction callback handler.
// If watched keys are modified before EXEC, the transaction will fail and return redis.TxFailedErr.
func (t *Transaction) Watch(ctx context.Context, handler func(tx *redis.Tx) error, keys ...string) error {
	if t.client == nil || t.client.UniversalClient == nil {
		return fmt.Errorf("redis client is uninitialized")
	}

	// Retrieve underlying redis client struct
	rClient, ok := t.client.UniversalClient.(*redis.Client)
	if !ok {
		return fmt.Errorf("universal client does not support optimistic transactions")
	}

	return rClient.Watch(ctx, handler, keys...)
}
