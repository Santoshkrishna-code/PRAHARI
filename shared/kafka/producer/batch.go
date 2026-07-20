package producer

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// PublishBatch writes multiple messages to the broker in a single transaction call.
func (p *Producer) PublishBatch(ctx context.Context, messages []kafka.Message) error {
	if p.writer == nil {
		return fmt.Errorf("producer writer is uninitialized")
	}

	err := p.writer.WriteMessages(ctx, messages...)
	if err != nil {
		return fmt.Errorf("failed to publish batch messages: %w", err)
	}

	return nil
}
