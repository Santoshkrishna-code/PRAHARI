package dlq

import (
	"context"
	"fmt"
)

// Handler processes routing failures and logs metadata.
type Handler struct {
	dlqProducer *DLQProducer
	dlqTopic    string
}

// NewHandler constructs a new Handler instance.
func NewHandler(dlqProducer *DLQProducer, dlqTopic string) *Handler {
	return &Handler{
		dlqProducer: dlqProducer,
		dlqTopic:    dlqTopic,
	}
}

// HandleFailure routes the failed payload to the DLQ.
func (h *Handler) HandleFailure(ctx context.Context, key, value []byte, processErr error) error {
	if h.dlqProducer == nil {
		return fmt.Errorf("dlq producer is uninitialized")
	}

	// Route to DLQ topic
	err := h.dlqProducer.RouteToDLQ(ctx, h.dlqTopic, key, value)
	if err != nil {
		return fmt.Errorf("failed to handle message failure: %w (original error: %v)", err, processErr)
	}

	return nil
}
