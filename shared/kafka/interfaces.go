package kafka

import (
	"context"
)

// Publisher defines the port for sending raw payloads to Kafka brokers.
type Publisher interface {
	Publish(ctx context.Context, topic string, key []byte, value []byte) error
}
