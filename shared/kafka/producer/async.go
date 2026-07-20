package producer

import (
	"context"
	"fmt"
)

// PublishAsync writes a single message asynchronously using background goroutines, returning immediately.
func (p *Producer) PublishAsync(ctx context.Context, topic string, key, value []byte) error {
	if p.writer == nil {
		return fmt.Errorf("producer writer is uninitialized")
	}

	// Dispatch write action to a background thread to prevent thread blocks
	go func() {
		// Use Background context to ensure lifecycle survives the calling request cancel
		_ = p.PublishSync(context.Background(), topic, key, value)
	}()

	return nil
}
