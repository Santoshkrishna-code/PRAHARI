package consumer

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// CommitMessage manually commits a message offset to the broker.
func (c *Consumer) CommitMessage(ctx context.Context, msg kafka.Message) error {
	if c.reader == nil {
		return fmt.Errorf("consumer reader is uninitialized")
	}
	return c.reader.CommitMessages(ctx, msg)
}
