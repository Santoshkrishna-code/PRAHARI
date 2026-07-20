package pubsub

import (
	"context"
	"fmt"

	prahariRedis "prahari/shared/redis"
)

// Subscriber manages Redis Pub/Sub channels subscriptions.
type Subscriber struct {
	client *prahariRedis.Client
}

// NewSubscriber constructs a new Subscriber.
func NewSubscriber(client *prahariRedis.Client) *Subscriber {
	return &Subscriber{client: client}
}

// Subscribe listens on a Redis channel and triggers the callback for each received payload.
// This is a blocking loop; launch in a goroutine to run asynchronously.
func (s *Subscriber) Subscribe(ctx context.Context, channel string, handler func(ctx context.Context, payload []byte) error) error {
	if s.client == nil || s.client.UniversalClient == nil {
		return fmt.Errorf("redis client is uninitialized")
	}

	pubsub := s.client.UniversalClient.Subscribe(ctx, channel)
	defer pubsub.Close()

	ch := pubsub.Channel()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg, ok := <-ch:
			if !ok {
				return nil
			}
			
			// Dispatch payload bytes to handler
			_ = handler(ctx, []byte(msg.Payload))
		}
	}
}
