package dlq_test

import (
	"context"
	"errors"
	"testing"

	"prahari/shared/kafka/dlq"
)

func TestDLQHandler_Uninitialized(t *testing.T) {
	// Constructing with nil dependencies returns proper validation errors
	h := dlq.NewHandler(nil, "test-dlq-topic")
	ctx := context.Background()

	err := h.HandleFailure(ctx, []byte("key"), []byte("value"), errors.New("original error"))
	if err == nil {
		t.Error("expected failure handling to error on nil DLQ producer, got nil")
	}
}
