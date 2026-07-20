package producer

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// PublishSync writes a single message to a topic synchronously and blocks until acknowledged.
func (p *Producer) PublishSync(ctx context.Context, topic string, key, value []byte) error {
	if p.writer == nil {
		return fmt.Errorf("producer writer is uninitialized")
	}

	msg := kafka.Message{
		Topic: topic,
		Key:   key,
		Value: value,
	}

	err := p.writer.WriteMessages(ctx, msg)
	if err != nil {
		return fmt.Errorf("failed to publish synchronous message: %w", err)
	}

	return nil
}
