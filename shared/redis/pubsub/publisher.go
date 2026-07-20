package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	prahariRedis "prahari/shared/redis"
)

// Publisher handles message publication to Redis Pub/Sub channels.
type Publisher struct {
	client *prahariRedis.Client
}

// NewPublisher constructs a new Publisher.
func NewPublisher(client *prahariRedis.Client) *Publisher {
	return &Publisher{client: client}
}

// Publish serializes and publishes a message to a specific Redis channel.
func (p *Publisher) Publish(ctx context.Context, channel string, message interface{}) error {
	if p.client == nil || p.client.UniversalClient == nil {
		return fmt.Errorf("redis client is uninitialized")
	}

	dataBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal pubsub message: %w", err)
	}

	err = p.client.UniversalClient.Publish(ctx, channel, dataBytes).Err()
	if err != nil {
		return fmt.Errorf("failed to publish message to channel %s: %w", channel, err)
	}

	return nil
}
