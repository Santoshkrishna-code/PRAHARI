package dlq

import (
	"context"
	"fmt"

	prahariProducer "prahari/shared/kafka/producer"
)

// DLQProducer handles writing poisoned event messages to dead-letter queue topics.
type DLQProducer struct {
	producer *prahariProducer.Producer
}

// NewDLQProducer constructs a new DLQProducer.
func NewDLQProducer(producer *prahariProducer.Producer) *DLQProducer {
	return &DLQProducer{producer: producer}
}

// RouteToDLQ sends the key/value payload to the specified DLQ topic.
func (d *DLQProducer) RouteToDLQ(ctx context.Context, dlqTopic string, key, value []byte) error {
	if d.producer == nil {
		return fmt.Errorf("underlying producer is uninitialized")
	}

	err := d.producer.PublishSync(ctx, dlqTopic, key, value)
	if err != nil {
		return fmt.Errorf("failed to route message to DLQ topic %s: %w", dlqTopic, err)
	}

	return nil
}
