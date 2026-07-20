package consumer_test

import (
	"context"
	"testing"

	prahariConsumer "prahari/shared/kafka/consumer"
)

func TestConsumer_Start_Uninitialized(t *testing.T) {
	ctx := context.Background()
	c := &prahariConsumer.Consumer{}

	err := c.StartConsumer(ctx, func(ctx context.Context, key, val []byte) error {
		return nil
	})
	if err == nil {
		t.Error("expected consumer loop to error on uninitialized reader, got nil")
	}
}
