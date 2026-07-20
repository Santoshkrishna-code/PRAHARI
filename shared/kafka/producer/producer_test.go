package producer_test

import (
	"context"
	"testing"

	prahariProducer "prahari/shared/kafka/producer"
)

func TestProducer_Publish_Uninitialized(t *testing.T) {
	// Constructing with empty config leaves writer uninitialized on mock,
	// or we can test error handling paths.
	ctx := context.Background()
	p := &prahariProducer.Producer{}

	err := p.PublishSync(ctx, "topic", []byte("key"), []byte("value"))
	if err == nil {
		t.Error("expected sync publish to error on uninitialized writer, got nil")
	}

	err = p.PublishAsync(ctx, "topic", []byte("key"), []byte("value"))
	if err == nil {
		t.Error("expected async publish to error on uninitialized writer, got nil")
	}
}
